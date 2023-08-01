package logic

import (
	"context"

	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

func (s *RoleRepository) FindUserRoles(userId int) (out []*entity2.Role, err error) {
	db := s.DbEngin
	var userRoles []entity2.UserRole
	err = db.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}

	var rids []int
	for _, item := range userRoles {
		rids = append(rids, item.RoleID)
	}

	var roles []*entity2.Role
	err = db.Where("id in (?)", rids).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// 获取Menu记录
func (s *RoleRepository) FindRoleMenus(roleId int) (list []*entity2.Menu, err error) {
	// 创建db
	db := s.DbEngin
	var roleMenus []*entity2.RoleMenu

	err = db.Where("role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, err
	}

	var menuIds []int
	for _, item := range roleMenus {
		menuIds = append(menuIds, item.MenuID)
	}

	var menus []*entity2.Menu
	err = db.Where("id in (?)", menus).Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

// 获取Api记录
func (s *RoleRepository) FindRoleApis(roleId int) (list []*entity2.Api, err error) {
	// 创建db
	db := s.DbEngin
	var roleApis []*entity2.RoleApi

	err = db.Where("role_id = ?", roleId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var apiIds []int
	for _, item := range roleApis {
		apiIds = append(apiIds, item.ApiID)
	}

	var apis []*entity2.Api
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
	var account entity2.UserAccount
	err = db.Where("user_id = ?", uid).First(&account).Error
	if err != nil {
		return nil, err
	}

	var userRoles []*entity2.UserRole
	for _, id := range roleIds {
		ur := &entity2.UserRole{
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
func (s *RoleRepository) UpdateRoleMenus(ctx context.Context, roleId int, menuIds []int) (role *entity2.Role, roleMenus []*entity2.RoleMenu, err error) {
	// 创建db
	db := s.DbEngin
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, nil, err
	}

	for _, id := range menuIds {
		rm := &entity2.RoleMenu{
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
func (s *RoleRepository) UpdateRoleResources(ctx context.Context, roleId int, apiIds []int) (role *entity2.Role, roleApis []*entity2.RoleApi, err error) {
	// 创建db
	db := s.DbEngin

	// 查询角色信息
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return role, nil, err
	}

	// 查询角色资源
	for _, id := range apiIds {
		ra := &entity2.RoleApi{
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
