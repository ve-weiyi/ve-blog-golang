package presencex_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/store/presencex"
	"github.com/ve-weiyi/blog-cloud/infra/store/storex"
)

// fakeMembers 是 storex.MemberStore 的内存实现：只保留存在性与各自的过期时刻，
// 与 RedisMemberStore 同语义（ttl <= 0 表示不过期）。
type fakeMembers struct {
	mu   sync.Mutex
	data map[string]map[string]time.Time // key -> member -> 过期时刻
}

func newFakeMembers() *fakeMembers {
	return &fakeMembers{data: make(map[string]map[string]time.Time)}
}

func (f *fakeMembers) Add(_ context.Context, key, member string, ttl time.Duration) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.data[key] == nil {
		f.data[key] = make(map[string]time.Time)
	}
	var expireAt time.Time
	if ttl > 0 {
		expireAt = time.Now().Add(ttl)
	}
	f.data[key][member] = expireAt

	return nil
}

func (f *fakeMembers) Members(ctx context.Context, key string) ([]string, error) {
	expiries, err := f.Expiries(ctx, key)
	if err != nil {
		return nil, err
	}

	members := make([]string, 0, len(expiries))
	for member := range expiries {
		members = append(members, member)
	}

	return members, nil
}

func (f *fakeMembers) Expiries(_ context.Context, key string) (map[string]time.Time, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	expiries := make(map[string]time.Time, len(f.data[key]))
	for member, expireAt := range f.data[key] {
		if expireAt.IsZero() || expireAt.After(time.Now()) {
			expiries[member] = expireAt
		}
	}

	return expiries, nil
}

func (f *fakeMembers) Has(_ context.Context, key, member string) (bool, error) {
	expiries, err := f.Expiries(context.Background(), key)
	if err != nil {
		return false, err
	}
	_, ok := expiries[member]

	return ok, nil
}

func (f *fakeMembers) Remove(_ context.Context, key string, members ...string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(members) == 0 {
		delete(f.data, key)
		return nil
	}
	for _, member := range members {
		delete(f.data[key], member)
	}

	return nil
}

var _ storex.MemberStore = (*fakeMembers)(nil)

func newTestStore(t *testing.T) *presencex.Store {
	t.Helper()

	store := presencex.NewStore(newFakeMembers(), "online", time.Minute)
	t.Cleanup(func() { _ = store.Close() })

	return store
}

// 成员的编解码是这套存储的对外格式契约：往返必须无损，缺分隔符的要能被认出来。
func TestMemberEncodeDecode(t *testing.T) {
	cases := []struct {
		name   string
		login  string
		ref    string
		wantOK bool
	}{
		{"普通登录", "uid-1", "1.2.3.4:5678", true},
		{"login 里带分隔符之外的字符", "a-b_c", "x", true},
		{"缺分隔符", "uid-1", "", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := presencex.EncodeMember(tc.login, tc.ref)

			login, ref, ok := presencex.DecodeMember(encoded)
			if ok != tc.wantOK {
				t.Fatalf("DecodeMember(%q) ok = %v, want %v", encoded, ok, tc.wantOK)
			}
			if !tc.wantOK {
				return
			}
			if login != tc.login || ref != tc.ref {
				t.Errorf("往返得到 (%q, %q), want (%q, %q)", login, ref, tc.login, tc.ref)
			}
		})
	}
}

// 在线状态按连接记：同一 login 的多条连接各算一条，两种口径因此不同。
func TestStore_CountsConnectionsNotUsers(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)

	for _, ref := range []string{"ref-a", "ref-b"} {
		if err := store.Add(ctx, "uid-1", ref, time.Minute); err != nil {
			t.Fatalf("Add(uid-1, %s): %v", ref, err)
		}
	}

	stats, err := store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumUsers != 1 || stats.NumClients != 2 {
		t.Errorf("stats = %+v, want {NumUsers:1 NumClients:2}", stats)
	}

	users, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 1 || users[0].Login != "uid-1" || users[0].Refs != 2 {
		t.Errorf("List = %+v, want 一条 uid-1 且 Refs 为 2", users)
	}

	if err := store.Remove(ctx, "uid-1", "ref-a"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	stats, err = store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumUsers != 1 || stats.NumClients != 1 {
		t.Errorf("断开其一后 stats = %+v, want {NumUsers:1 NumClients:1}", stats)
	}
}

// Add 是幂等的：同一 (login, ref) 重复上报（续期）不该多出一条连接。
func TestStore_AddIsIdempotent(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)

	for i := 0; i < 3; i++ {
		if err := store.Add(ctx, "uid-1", "ref-a", time.Minute); err != nil {
			t.Fatalf("Add 第 %d 次: %v", i+1, err)
		}
	}

	stats, err := store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumClients != 1 {
		t.Errorf("重复上报后 NumClients = %d, want 1", stats.NumClients)
	}
}

// 按用户聚合时取最新的续期时刻。
func TestStore_ListTakesLatestRenewal(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)

	if err := store.Add(ctx, "uid-1", "ref-old", time.Minute); err != nil {
		t.Fatalf("Add(ref-old): %v", err)
	}
	time.Sleep(10 * time.Millisecond)
	if err := store.Add(ctx, "uid-1", "ref-new", time.Minute); err != nil {
		t.Fatalf("Add(ref-new): %v", err)
	}

	users, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("List = %+v, want 一条", users)
	}
	if users[0].LastActiveAt < time.Now().Add(-time.Second).UnixMilli() {
		t.Errorf("LastActiveAt = %d, want 接近当前时刻", users[0].LastActiveAt)
	}
	if users[0].Refs != 2 {
		t.Errorf("Refs = %d, want 2", users[0].Refs)
	}
}

// 存活判定来自成员的过期时刻：过了 TTL 就既不在名单里、也不参与计数。
func TestStore_ExpiredConnectionDisappears(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)

	if err := store.Add(ctx, "uid-1", "ref-a", 20*time.Millisecond); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if err := store.Add(ctx, "uid-2", "ref-b", time.Hour); err != nil {
		t.Fatalf("Add: %v", err)
	}

	time.Sleep(80 * time.Millisecond)

	stats, err := store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumUsers != 1 || stats.NumClients != 1 {
		t.Errorf("超期后 stats = %+v, want 只剩 uid-2", stats)
	}

	users, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 1 || users[0].Login != "uid-2" {
		t.Errorf("超期后 List = %+v, want 只剩 uid-2", users)
	}
}

// 键里混进没有分隔符的成员（脏数据）时，读取要跳过它而不是算成一条连接。
func TestStore_IgnoresMalformedMember(t *testing.T) {
	ctx := context.Background()
	store := newTestStore(t)

	if err := store.Add(ctx, "uid-1", "ref-a", time.Minute); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if err := store.Add(ctx, "garbage", "", 0); err != nil {
		t.Fatalf("Add: %v", err)
	}

	stats, err := store.PresenceStats(ctx)
	if err != nil {
		t.Fatalf("PresenceStats: %v", err)
	}
	if stats.NumUsers != 1 || stats.NumClients != 1 {
		t.Errorf("stats = %+v, want 只算 uid-1", stats)
	}
}
