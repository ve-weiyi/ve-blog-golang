package logic

import (
	"context"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
)

func (s *RoleRepository) FindUserRoles(userId int) (out []*entity.Role, err error) {
	db := s.DbEngin
	var userRoles []entity.UserRole
	err = db.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}

	var rids []int
	for _, item := range userRoles {
		rids = append(rids, item.RoleID)
	}

	var roles []*entity.Role
	err = db.Where("id in (?)", rids).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// 获取Menu记录
func (s *RoleRepository) FindRoleMenus(roleId int) (list []*entity.Menu, err error) {
	// 创建db
	db := s.DbEngin
	var roleMenus []*entity.RoleMenu

	err = db.Where("role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, err
	}

	var menuIds []int
	for _, item := range roleMenus {
		menuIds = append(menuIds, item.MenuID)
	}

	var menus []*entity.Menu
	err = db.Where("id in (?)", menus).Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

// 获取Api记录
func (s *RoleRepository) FindRoleApis(roleId int) (list []*entity.Api, err error) {
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

// 修改用户角色
func (s *RoleRepository) UpdateUserRoles(ctx context.Context, uid int, roleIds []int) (data interface{}, err error) {
	// 创建db
	db := s.DbEngin
	var account entity.UserAccount
	err = db.Where("user_id = ?", uid).First(&account).Error
	if err != nil {
		return nil, err
	}

	var userRoles []*entity.UserRole
	for _, id := range roleIds {
		ur := &entity.UserRole{
			UserID: account.ID,
			RoleID: id,
		}
		userRoles = append(userRoles, ur)
	}

	// 开启事务
	tx := db.Begin()

	//先删除所有菜单，再添加
	err = tx.Delete(&userRoles, "user_id = ?", uid).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//再添加
	err = tx.Create(&userRoles).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//提交事务
	tx.Commit()

	return nil, err
}

// 设置角色菜单
func (s *RoleRepository) UpdateRoleMenus(ctx context.Context, roleId int, menuIds []int) (role *entity.Role, roleMenus []*entity.RoleMenu, err error) {
	// 创建db
	db := s.DbEngin
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, nil, err
	}

	for _, id := range menuIds {
		rm := &entity.RoleMenu{
			RoleID: role.ID,
			MenuID: id,
		}
		roleMenus = append(roleMenus, rm)
	}

	// 开启事务
	tx := db.Begin()

	//先删除所有菜单，再添加
	err = tx.Delete(&roleMenus, "role_id = ?", role.ID).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	//再添加
	err = tx.Create(&roleMenus).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	//提交事务
	tx.Commit()

	return role, roleMenus, err
}

// 设置角色菜单
func (s *RoleRepository) UpdateRoleResources(ctx context.Context, roleId int, apiIds []int) (role *entity.Role, roleApis []*entity.RoleApi, err error) {
	// 创建db
	db := s.DbEngin

	// 查询角色信息
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return role, nil, err
	}

	// 查询角色资源
	for _, id := range apiIds {
		ra := &entity.RoleApi{
			RoleID: role.ID,
			ApiID:  id,
		}
		roleApis = append(roleApis, ra)
	}

	// 开启事务
	tx := db.Begin()

	//先删除所有菜单，再添加
	err = tx.Delete(&roleApis, "role_id = ?", role.ID).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	//再添加
	err = tx.Create(&roleApis).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	//提交事务
	tx.Commit()

	return role, roleApis, err
}
