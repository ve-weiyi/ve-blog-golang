package rbac

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

// 删除CasbinRule记录
func (s *CachedEnforcer) DeleteRolePolicy(roleName string, domain string) (result interface{}, err error) {
	rbac := s.SyncedCachedEnforcer
	return rbac.RemoveFilteredNamedPolicy("p", 0, roleName, domain)
}

// 添加CasbinRule记录
func (s *CachedEnforcer) AddRolePolicy(roleName string, domain string, apis []*entity.Api) (result interface{}, err error) {
	var policies [][]string
	for _, item := range apis {
		if item.Path == "" || item.Method == "" {
			continue
		}
		policies = append(policies, []string{roleName, domain, item.Path, item.Method})

	}

	rbac := s.SyncedCachedEnforcer
	_, err = rbac.AddPolicies(policies)
	if err != nil {
		return nil, err
	}
	return nil, rbac.SavePolicy()
}
