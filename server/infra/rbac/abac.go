package rbac

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type Enforcer interface {
	LoadPermissions() error
	VerifyUserPermissions(uid int, path string, method string) (bool, error)
}

type ApiPermission struct {
	entity.Api
	Roles []*entity.Role
}

// 判断角色拥有哪些资源（Resource-Based Access Control）
type ResourceEnforcer struct {
	DbEngin    *gorm.DB
	cacheEngin *redis.Client
	// 资源角色
	apiMap map[string]ApiPermission
}

func NewResourceEnforcer(db *gorm.DB, rdb *redis.Client) *ResourceEnforcer {
	return &ResourceEnforcer{
		DbEngin:    db,
		cacheEngin: rdb,
		apiMap:     make(map[string]ApiPermission),
	}
}

// 从数据库加载角色的资源
func (s *ResourceEnforcer) LoadPermissions() error {
	// 查询所有资源
	var apis []entity.Api
	err := s.DbEngin.Find(&apis).Error
	if err != nil {
		return err
	}

	// 查询资源所需角色
	for _, api := range apis {
		roles, err := s.findApiRoles(api.ID)
		if err != nil {
			return err
		}

		ap := ApiPermission{
			Api:   api,
			Roles: roles,
		}
		s.apiMap[fmt.Sprintf("%s [%s]", api.Path, api.Method)] = ap
	}
	return nil
}

func (s *ResourceEnforcer) VerifyUserPermissions(uid int, path string, method string) (bool, error) {
	// 判断该资源所需的角色是否包含该角色
	api, exists := s.apiMap[fmt.Sprintf("%s [%s]", path, method)]
	if !exists {
		log.Println("VerifyUserPermissions fail:api not found")
		return true, nil
	}

	// 关闭了当前接口
	if api.Status == 0 {
		return false, fmt.Errorf("VerifyUserPermissions fail:api is closed")
	}

	// 查询用户角色
	var urs []entity.UserRole
	err := s.DbEngin.Where("user_id = ?", uid).Find(&urs).Error
	// 用户角色为空，返回false
	if err != nil {
		return false, fmt.Errorf("VerifyUserPermissions fail:%v", err)
	}

	// 遍历资源角色
	for _, ur := range api.Roles {
		// 匹配用户角色
		for _, u := range urs {
			if u.RoleID == ur.ID {
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("VerifyUserPermissions fail:permissions not match")
}

func (s *ResourceEnforcer) GetApiPermission(path string, method string) *ApiPermission {
	// 判断该资源所需的角色是否包含该角色
	api, exists := s.apiMap[fmt.Sprintf("%s [%s]", path, method)]
	if !exists {
		return nil
	}

	return &api
}

// 获取Api记录
func (s *ResourceEnforcer) findApiRoles(apiId int) (list []*entity.Role, err error) {
	ctx := context.Background()
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	var roleApis []*entity.RoleApi

	err = db.Where("api_id = ?", apiId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var roleIds []int
	for _, item := range roleApis {
		roleIds = append(roleIds, item.ApiID)
	}

	var roles []*entity.Role
	err = db.Where("id in (?)", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
