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
type RbacHolder struct {
	sync.RWMutex

	pr permissionrpc.PermissionRpc

	// key: (user), value: role_name[]
	lru *collection.Cache

	// key: (role_name), value: role
	roles map[string]*permissionrpc.RoleDetails

	// key: (api_name), value: api
	apis map[string]*permissionrpc.ApiDetails

	// key: roleId, value: apiIds
	rs map[int64][]int64
}

func NewRbacHolder(pr permissionrpc.PermissionRpc) *RbacHolder {
	lru, err := collection.NewCache(30 * time.Minute)
	if err != nil {
		panic(err)
	}

	return &RbacHolder{
		pr:  pr,
		lru: lru,
	}
}

func (m *RbacHolder) LoadPolicy() error {
	m.Lock()
	defer m.Unlock()

	var rs = make(map[int64][]int64)
	var roles = make(map[string]*permissionrpc.RoleDetails)
	var apis = make(map[string]*permissionrpc.ApiDetails)

	// 收集角色
	roleList, err := m.pr.FindRoleList(context.Background(), &permissionrpc.FindRoleListReq{})
	if err != nil {
		return err
	}

	for _, v := range roleList.List {
		resource, err := m.pr.FindRoleResources(context.Background(), &permissionrpc.IdReq{Id: v.Id})
		if err != nil {
			return err
		}

		roles[v.RoleName] = v
		rs[v.Id] = resource.ApiIds
	}

	apiList, err := m.pr.FindAllApi(context.Background(), &permissionrpc.EmptyReq{})
	if err != nil {
		return err
	}
	for _, v := range apiList.List {
		apis[getResourceKey(v.Path, v.Method)] = v
	}

	m.roles = roles
	m.apis = apis
	m.rs = rs
	return nil
}

func (m *RbacHolder) Enforce(ctx context.Context, user string, resource string, action string) error {
	m.RLock()
	defer m.RUnlock()

	// 获取资源
	api, err := m.GetApi(resource, action)
	if err != nil {
		return err
	}

	// 获取用户角色
	userRoles, err := m.getUserRoles(user)
	if err != nil {
		return err
	}

	for _, roleId := range userRoles {
		// 获取角色资源
		role, err := m.GetRole(roleId)
		if err != nil {
			return err
		}

		// 判断用户是否有权限
		if apiIds, ok := m.rs[role.Id]; ok {
			for _, apiId := range apiIds {
				if apiId == api.Id {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("user %s has no permission to access %s %s", user, resource, action)
}

func (m *RbacHolder) getUserRoles(userId string) ([]string, error) {
	// 从缓存中获取用户角色
	roles, ok := m.lru.Get(userId)
	if ok {
		return roles.([]string), nil
	}

	// 从数据库查找
	urs, err := m.pr.FindUserRoles(context.Background(), &permissionrpc.UserIdReq{UserId: userId})
	if err != nil {
		return nil, err
	}

	var roleIds []string
	for _, v := range urs.List {
		roleIds = append(roleIds, v.RoleName)
	}

	// 将用户角色放入缓存
	m.lru.Set(userId, roleIds)
	return roleIds, nil
}

func (m *RbacHolder) GetApi(path string, method string) (*permissionrpc.ApiDetails, error) {
	key := getResourceKey(path, method)

	api, ok := m.apis[key]
	if !ok {
		return nil, fmt.Errorf("resource %s not found", key)
	}

	if api.IsDisable == 1 {
		return nil, fmt.Errorf("resource %s is disabled", key)
	}

	return api, nil
}

func (m *RbacHolder) GetRole(key string) (*permissionrpc.RoleDetails, error) {
	role, ok := m.roles[key]
	if !ok {
		return nil, fmt.Errorf("role %s not found", key)
	}

	if role.IsDisable == 1 {
		return nil, fmt.Errorf("role %s is disabled", key)
	}

	return role, nil
}

func getResourceKey(path string, method string) string {
	return fmt.Sprintf("%s %s", path, method)
}
