package permissionx

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/mikespook/gorbac/v2"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/accessservice"
	"github.com/ve-weiyi/blog-cloud/infra/domain/cachekey"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

// registered 是「已注册接口」全集，用于填充登记索引。
// Enforce 对未登记的接口 fail-closed，因此被拒绝的用例必须先把接口登记进来。
func newTestEnforcer(userRoles map[string][]string, rolePerms map[string][]string, registered []string) *RbacEnforcer {
	rbac := gorbac.New()
	for role, perms := range rolePerms {
		r := gorbac.NewStdRole(role)
		for _, perm := range perms {
			_ = r.Assign(apiPermission(perm))
		}
		_ = rbac.Add(r)
	}

	allPerms := make(map[string]struct{}, len(registered))
	registeredByMethod := make(map[string][]string)
	for _, perm := range registered {
		allPerms[perm] = struct{}{}
		method, path, ok := splitPermID(perm)
		if ok {
			registeredByMethod[method] = append(registeredByMethod[method], normalizePath(path))
		}
	}

	cache := make(map[string]userRoleEntry, len(userRoles))
	for user, roles := range userRoles {
		cache[user] = userRoleEntry{roles: roles, loadedAt: time.Now()}
	}

	return &RbacEnforcer{
		rbac:               rbac,
		userRoles:          cache,
		allPerms:           allPerms,
		registeredByMethod: registeredByMethod,
		userRoleMaxAge:     time.Minute, // 测试期间不应过期；零值会让内存缓存立即失效
		policyLoaded:       true,
	}
}

// fakeAccessService 只覆写 LoadPolicy / getUserRoles 用得到的方法；
// 其余靠嵌入接口兜底（一旦被用到即 nil panic，测试会立刻暴露）
type fakeAccessService struct {
	accessservice.AccessService
	roles           []*accessservice.Role
	apis            []*accessservice.Api
	resources       map[int64][]int64
	userRoles       []*accessservice.Role
	onListUserRoles func()
}

func (f *fakeAccessService) ListRoles(context.Context, *accessservice.ListRolesRequest, ...grpc.CallOption) (*accessservice.ListRolesResponse, error) {
	return &accessservice.ListRolesResponse{List: f.roles}, nil
}

func (f *fakeAccessService) ListApis(context.Context, *accessservice.ListApisRequest, ...grpc.CallOption) (*accessservice.ListApisResponse, error) {
	return &accessservice.ListApisResponse{List: f.apis}, nil
}

func (f *fakeAccessService) GetRoleResource(_ context.Context, in *accessservice.GetRoleResourceRequest, _ ...grpc.CallOption) (*accessservice.GetRoleResourceResponse, error) {
	return &accessservice.GetRoleResourceResponse{ApiIds: f.resources[in.RoleId]}, nil
}

func (f *fakeAccessService) GetRoleResources(_ context.Context, in *accessservice.GetRoleResourcesRequest, _ ...grpc.CallOption) (*accessservice.GetRoleResourcesResponse, error) {
	list := make([]*accessservice.RoleResource, 0, len(in.RoleIds))
	for _, roleId := range in.RoleIds {
		list = append(list, &accessservice.RoleResource{
			RoleId: roleId,
			ApiIds: f.resources[roleId],
		})
	}
	return &accessservice.GetRoleResourcesResponse{List: list}, nil
}

func (f *fakeAccessService) ListUserRoles(context.Context, *accessservice.ListUserRolesRequest, ...grpc.CallOption) (*accessservice.ListUserRolesResponse, error) {
	if f.onListUserRoles != nil {
		f.onListUserRoles()
	}
	return &accessservice.ListUserRolesResponse{List: f.userRoles}, nil
}

// 禁用一个角色，不应把它独占的接口从「已注册」集合里摘掉。
//
// Enforce 对未注册的接口是直接放行的，所以接口一旦因为角色被禁用而"注销"，
// 就会从「仅该角色可访问」翻转为「对所有人放行」——与禁用角色的预期恰好相反。
func TestLoadPolicyKeepsDisabledRoleApisRegistered(t *testing.T) {
	svc := &fakeAccessService{
		roles: []*accessservice.Role{
			{Id: 1, RoleKey: "admin", Status: enums.RoleStatusNormal},
			{Id: 2, RoleKey: "auditor", Status: enums.RoleStatusDisabled},
		},
		apis: []*accessservice.Api{
			{Id: 10, Method: "GET", Path: "/admin-api/v1/audit"},
			{Id: 11, Method: "GET", Path: "/admin-api/v1/users"},
		},
		resources: map[int64][]int64{
			1: {11},
			2: {10}, // 只分配给已禁用的角色
		},
	}

	e := &RbacEnforcer{pr: svc, rbac: gorbac.New(), userRoles: make(map[string]userRoleEntry)}
	if err := e.LoadPolicy(); err != nil {
		t.Fatalf("LoadPolicy: %v", err)
	}

	if _, ok := e.allPerms["GET:/admin-api/v1/audit"]; !ok {
		t.Fatal("被禁用角色独占的接口从登记集合中消失了——它会变成对所有人放行")
	}

	// 登记了，但不该真的被授权：已禁用的角色不进授权图
	if ok, _ := e.Enforce("someone", "/admin-api/v1/audit", "GET"); ok {
		t.Fatal("被禁用角色不应授予任何权限")
	}
}

func TestEnforce(t *testing.T) {
	e := newTestEnforcer(
		map[string][]string{
			"user1": {"admin"},
			"user2": {"viewer", "editor"},
			"root":  {"root"},
		},
		map[string][]string{
			"admin":  {"GET:/api/v1/users", "GET:/api/v1/users/:id"},
			"viewer": {"GET:/api/v1/posts", "GET:/api/v1/orders/{id}"},
			"editor": {"POST:/api/v1/posts"},
		},
		[]string{
			"GET:/api/v1/users", "GET:/api/v1/users/:id",
			"GET:/api/v1/posts", "POST:/api/v1/posts",
			"GET:/api/v1/orders/{id}",
			// 已注册但未授予任何角色 —— 覆盖"接口存在但无权访问"的拒绝路径
			"POST:/api/v1/users", "DELETE:/api/v1/posts",
		},
	)

	cases := []struct {
		name    string
		user    string
		path    string
		method  string
		wantErr bool
	}{
		{"allowed", "user1", "/api/v1/users", "GET", false},
		{"denied", "user1", "/api/v1/users", "POST", true},
		{"root bypass", "root", "/api/v1/anything", "DELETE", false},
		{"param path", "user1", "/api/v1/users/123", "GET", false},
		{"param path denied", "user1", "/api/v1/users/123", "DELETE", true},
		{"brace param", "user2", "/api/v1/orders/42", "GET", false},
		{"unregistered denied", "user1", "/api/v1/unknown", "GET", true},
		{"multiple roles allowed", "user2", "/api/v1/posts", "POST", false},
		{"multiple roles denied", "user2", "/api/v1/posts", "DELETE", true},
		{"empty user", "", "/api/v1/users", "GET", true},
		{"empty path", "user1", "", "GET", true},
		{"empty method", "user1", "/api/v1/users", "", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := e.Enforce(c.user, c.path, c.method)
			if (err != nil) != c.wantErr {
				t.Errorf("Enforce(%q, %q, %q) error=%v, wantErr=%v", c.user, c.path, c.method, err, c.wantErr)
			}
		})
	}
}

func newTestEnforcerWithRedis(t *testing.T, pr accessservice.AccessService) (*RbacEnforcer, *miniredis.Miniredis) {
	t.Helper()

	mr := miniredis.RunT(t)
	rds := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rds.Close() })

	return &RbacEnforcer{
		pr:             pr,
		rds:            rds,
		userRoles:      make(map[string]userRoleEntry),
		userRoleGen:    make(map[string]uint64),
		userRoleTTL:    5 * time.Minute,
		userRoleMaxAge: 5 * time.Minute,
	}, mr
}

// 撤权发生在 RPC 返回之后、结果写入缓存之前：本次请求可以继续用拉取结果，
// 但结果不得写入内存或 Redis，否则撤权会在 TTL 内失效。
func TestGetUserRoles_DiscardsResultWhenInvalidatedInFlight(t *testing.T) {
	pr := &fakeAccessService{userRoles: []*accessservice.Role{{Id: 1, RoleKey: "admin"}}}
	e, mr := newTestEnforcerWithRedis(t, pr)
	pr.onListUserRoles = func() { e.InvalidateUser("u1") }

	roles, err := e.getUserRoles("u1")
	if err != nil {
		t.Fatalf("getUserRoles: %v", err)
	}
	if len(roles) != 1 {
		t.Fatalf("roles = %v, want 1 个", roles)
	}

	if mr.Exists(cachekey.UserRoleCacheKey("u1")) {
		t.Error("在途结果被写回了 Redis 缓存")
	}
	e.mu.RLock()
	_, ok := e.userRoles["u1"]
	e.mu.RUnlock()
	if ok {
		t.Error("在途结果被写回了内存缓存")
	}
}

// Redis 中的缓存代际落后于当前代际时不可信，必须回源并写入新代际。
func TestGetUserRoles_RejectsStaleGeneration(t *testing.T) {
	pr := &fakeAccessService{userRoles: []*accessservice.Role{{Id: 1, RoleKey: "admin"}}}
	e, mr := newTestEnforcerWithRedis(t, pr)

	if err := mr.Set(cachekey.UserRoleCacheKey("u1"), `{"generation":0,"roles":["stale:1"]}`); err != nil {
		t.Fatalf("set stale cache: %v", err)
	}
	if err := mr.Set(cachekey.UserRoleGenKey("u1"), "1"); err != nil {
		t.Fatalf("set generation: %v", err)
	}

	roles, err := e.getUserRoles("u1")
	if err != nil {
		t.Fatalf("getUserRoles: %v", err)
	}
	if len(roles) != 1 || roles[0] == "stale:1" {
		t.Fatalf("roles = %v, want 回源结果", roles)
	}

	val, err := mr.Get(cachekey.UserRoleCacheKey("u1"))
	if err != nil {
		t.Fatalf("get cache: %v", err)
	}
	if !strings.Contains(val, `"generation":1`) {
		t.Errorf("cache = %s, want generation 1", val)
	}
}
