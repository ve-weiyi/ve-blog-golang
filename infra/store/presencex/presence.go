package presencex

import (
	"sort"
	"strings"
	"time"
)

// memberSep 分隔成员里的 login 与 ref。login（用户 UUID 或游客设备号）与
// ref（连接的 clientId）都不会含 NUL。
const memberSep = "\x00"

// DefaultTTL 是在线条目的默认存活时长：续期停了多久算离线。
const DefaultTTL = 60 * time.Second

// EncodeMember 把一条连接编成成员名。这份格式是**跨服务的数据契约** ——
// 写的一方与读的一方都只经由这里，格式变更因此只有一处。
func EncodeMember(login, ref string) string {
	return login + memberSep + ref
}

// DecodeMember 拆回登录标识与连接标识；缺任一半都不算一条合法连接。
func DecodeMember(member string) (login, ref string, ok bool) {
	login, ref, ok = strings.Cut(member, memberSep)
	return login, ref, ok && login != "" && ref != ""
}

// OnlineUser 是消费方视角的一条在线用户：同一 login 的多条连接合成一条。
type OnlineUser struct {
	Login        string
	LastActiveAt int64 // 最近一次续期的时刻（unix 毫秒）
	Refs         int   // 当前连接数
}

// aggregate 把成员聚合成用户：按 login 合并，取最新的续期时刻。
// 拆不出 login/ref 的成员跳过 —— 脏数据不该被算成一条在线连接。
func aggregate(expiries map[string]time.Time, ttl time.Duration) []OnlineUser {
	byLogin := make(map[string]*OnlineUser, len(expiries))
	for member, expireAt := range expiries {
		login, _, ok := DecodeMember(member)
		if !ok {
			continue
		}

		at := expireToLastActive(expireAt, ttl)
		user, seen := byLogin[login]
		if !seen {
			byLogin[login] = &OnlineUser{Login: login, LastActiveAt: at, Refs: 1}
			continue
		}
		user.Refs++
		if at > user.LastActiveAt {
			user.LastActiveAt = at
		}
	}

	users := make([]OnlineUser, 0, len(byLogin))
	for _, user := range byLogin {
		users = append(users, *user)
	}
	sort.Slice(users, func(i, j int) bool { return users[i].Login < users[j].Login })

	return users
}

// expireToLastActive 由成员的过期时刻反推续期时刻：过期 = 续期 + 存活时长。
// 不过期（零值时刻）的成员没有续期时刻可言，给 0。
func expireToLastActive(expireAt time.Time, ttl time.Duration) int64 {
	if expireAt.IsZero() {
		return 0
	}

	return expireAt.Add(-ttl).UnixMilli()
}
