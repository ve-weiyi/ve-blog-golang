package repository

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

func (s *RoleRepository) FindUserRoles(ctx context.Context, userId int64) (out []*entity.Role, err error) {
	db := s.DbEngin.WithContext(ctx)
	var userRoles []entity.UserRole
	err = db.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}

	var rids []int64
	for _, item := range userRoles {
		rids = append(rids, item.RoleId)
	}

	var roles []*entity.Role
	err = db.Where("id in (?)", rids).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// 获取Menu记录
func (s *RoleRepository) FindRoleMenus(ctx context.Context, roleId int64) (list []*entity.Menu, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	var roleMenus []*entity.RoleMenu

	err = db.Where("role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, err
	}

	var menuIds []int64
	for _, item := range roleMenus {
		menuIds = append(menuIds, item.MenuId)
	}

	var menus []*entity.Menu
	err = db.Where("id in (?)", menus).Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

// 获取Api记录
func (s *RoleRepository) FindRoleApis(ctx context.Context, roleId int64) (list []*entity.Api, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	var roleApis []*entity.RoleApi

	err = db.Where("role_id = ?", roleId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var apiIds []int64
	for _, item := range roleApis {
		apiIds = append(apiIds, item.ApiId)
	}

	var apis []*entity.Api
	err = db.Where("id in (?)", apiIds).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}

// 修改用户角色
func (s *RoleRepository) UpdateUserRoles(ctx context.Context, uid int64, roleIds []int64) (data interface{}, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	var account entity.UserAccount
	err = db.Where("id = ?", uid).First(&account).Error
	if err != nil {
		return nil, err
	}

	var userRoles []*entity.UserRole
	for _, id := range roleIds {
		ur := &entity.UserRole{
			UserId: account.Id,
			RoleId: id,
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
func (s *RoleRepository) UpdateRoleMenus(ctx context.Context, roleId int64, menuIds []int64) (role *entity.Role, count int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, 0, err
	}

	// 开启事务
	tx := db.Begin()

	//先删除所有菜单，再添加
	err = tx.Delete(&entity.RoleMenu{}, "role_id = ?", role.Id).Error
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	var roleMenus []*entity.RoleMenu
	//再添加
	for _, id := range menuIds {
		rm := &entity.RoleMenu{
			RoleId: role.Id,
			MenuId: id,
		}
		roleMenus = append(roleMenus, rm)
	}

	err = tx.CreateInBatches(&roleMenus, len(roleMenus)).Error
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	count++
	//提交事务
	tx.Commit()

	return role, count, err
}

// 设置角色菜单
func (s *RoleRepository) UpdateRoleApis(ctx context.Context, roleId int64, apiIds []int64) (role *entity.Role, count int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 查询角色信息
	err = db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, 0, err
	}

	// 开启事务
	tx := db.Begin()

	//先删除所有菜单，再添加
	err = tx.Delete(&entity.RoleApi{}, "role_id = ?", role.Id).Error
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	//再添加
	var roleApis []*entity.RoleApi
	for _, id := range apiIds {
		ra := &entity.RoleApi{
			RoleId: role.Id,
			ApiId:  id,
		}
		roleApis = append(roleApis, ra)
	}

	err = tx.CreateInBatches(&roleApis, len(roleApis)).Error
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	//提交事务
	tx.Commit()

	return role, count, err
}
