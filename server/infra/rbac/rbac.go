package rbac

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

type RolePermission struct {
	entity.Role
	Apis []*entity.Api
}

// 判断角色拥有哪些资源（Role-Based Access Control，RBAC）
type RoleEnforcer struct {
	DbEngin *gorm.DB
	// 角色权限
	roleMap map[int]RolePermission
}

// 从数据库加载角色的资源
func (s *RoleEnforcer) LoadPermissions() error {

	// 查询所有角色
	var roles []entity.Role
	err := s.DbEngin.Find(&roles).Error
	if err != nil {
		return err
	}

	// 查询角色权限
	for _, role := range roles {
		apis, err := s.FindRoleApis(role.ID)
		if err != nil {
			return err
		}

		rp := RolePermission{
			Role: role,
			Apis: apis,
		}
		s.roleMap[role.ID] = rp
	}
	return nil
}

func (s *RoleEnforcer) VerifyUserPermissions(uid int, path string, method string) (bool, error) {
	// 查询用户角色
	var urs []entity.UserRole
	s.DbEngin.Where("user_id = ?", uid).Find(&urs)
	if len(urs) == 0 {
		return false, fmt.Errorf("VerifyUserPermissions fail:user not found")
	}

	// 判断用户角色是否包含该资源
	for _, ur := range urs {
		role, exist := s.roleMap[ur.RoleID]
		if !exist {
			continue
		}
		// 匹配用户角色
		for _, api := range role.Apis {
			if api.Path == path && api.Method == method {
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("VerifyUserPermissions fail:permissions not match")
}

// 获取Api记录
func (s *RoleEnforcer) FindRoleApis(roleId int) (list []*entity.Api, err error) {
	// 创建db
	db := s.DbEngin
	var roleApis []*entity.RoleApi

	err = db.Where("role_id = ?", roleId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var apiIds []int
	for _, item := range roleApis {
		apiIds = append(apiIds, item.ApiID)
	}

	var apis []*entity.Api
	err = db.Where("id in (?)", apiIds).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}
