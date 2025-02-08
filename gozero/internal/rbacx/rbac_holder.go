package rbacx

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/collection"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/permissionrpc"
)

// 角色资源管理器
type MemoryHolder struct {
	sync.RWMutex

	pr permissionrpc.PermissionRpc

	autoload bool

	// key: (user), value: role_name[]
	lru *collection.Cache

	// 用户角色缓存 key: (user), value: role_name[]
	user map[string][]string

	// key: roleId, value: apiIds
	policy map[string][]string

	// key: (role_name), value: role
	roles map[string]*permissionrpc.RoleDetails

	// key: (api_name), value: api
	apis map[string]*permissionrpc.ApiDetails
}

func NewMemoryHolder(pr permissionrpc.PermissionRpc) *MemoryHolder {
	lru, err := collection.NewCache(30 * time.Minute)
	if err != nil {
		panic(err)
	}

	return &MemoryHolder{
		pr:       pr,
		lru:      lru,
		autoload: false,
	}
}

func (m *MemoryHolder) StartAutoLoadPolicy() {
	m.autoload = true
}

func (m *MemoryHolder) ClearPolicy() error {
	m.Lock()
	defer m.Unlock()

	m.policy = make(map[string][]string)
	m.user = make(map[string][]string)
	m.roles = make(map[string]*permissionrpc.RoleDetails)
	m.apis = make(map[string]*permissionrpc.ApiDetails)

	if m.autoload {
		return m.LoadPolicy()
	}
	return nil
}

func (m *MemoryHolder) LoadPolicy() error {
	m.Lock()
	defer m.Unlock()

	var rs = make(map[int64][]int64)
	var roles = make(map[int64]*permissionrpc.RoleDetails)
	var apis = make(map[int64]*permissionrpc.ApiDetails)

	// 收集角色
	roleList, err := m.pr.FindRoleList(context.Background(), &permissionrpc.FindRoleListReq{})
	if err != nil {
		return err
	}
	for _, v := range roleList.List {
		roles[v.Id] = v
	}

	// 收集资源
	apiList, err := m.pr.FindAllApi(context.Background(), &permissionrpc.EmptyReq{})
	if err != nil {
		return err
	}
	for _, v := range apiList.List {
		apis[v.Id] = v
	}

	// 收集角色-资源
	for _, v := range roleList.List {
		resource, err := m.pr.FindRoleResources(context.Background(), &permissionrpc.IdReq{Id: v.Id})
		if err != nil {
			return err
		}

		rs[v.Id] = resource.ApiIds
	}

	for roleId, apiIds := range rs {
		role, ok := roles[roleId]
		if !ok {
			return fmt.Errorf("role %d not found", roleId)
		}

		if role.IsDisable == 1 {
			continue
		}

		for _, apiId := range apiIds {
			api, ok := apis[apiId]
			if !ok {
				return fmt.Errorf("api %d not found", apiId)
			}

			if api.IsDisable == 1 {
				continue
			}

			m.policy[role.RoleName] = append(m.policy[role.RoleName], fmt.Sprintf("%s %s", api.Path, api.Method))
		}
	}

	return nil
}

func (m *MemoryHolder) Enforce(user string, resource string, action string) error {
	m.RLock()
	defer m.RUnlock()

	err := m.dynamicLoadUserRoles(user)
	if err != nil {
		return err
	}

	ok, err := m.enforce(user, resource, action)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("permission denied, user: %s, resource: %s, action: %s", user, resource, action)
	}

	return nil
}

// 加载用户角色
func (m *MemoryHolder) dynamicLoadUserRoles(userId string) error {
	if _, ok := m.user[userId]; ok {
		return nil
	}

	// 获取用户角色
	userRoles, err := m.pr.FindUserRoles(context.Background(), &permissionrpc.UserIdReq{UserId: userId})
	if err != nil {
		return err
	}

	for _, v := range userRoles.List {
		m.user[userId] = append(m.user[userId], v.RoleName)
	}

	return nil
}

func (m *MemoryHolder) enforce(user string, resource string, action string) (bool, error) {
	m.RLock()
	defer m.RUnlock()

	roles, ok := m.user[user]
	if !ok {
		return false, nil
	}

	p := fmt.Sprintf("%s %s", resource, action)
	for _, role := range roles {
		policies, ok := m.policy[role]
		if !ok {
			continue
		}

		for _, policy := range policies {
			if policy == p {
				return true, nil
			}
		}
	}

	return false, nil
}
