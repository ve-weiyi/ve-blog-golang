// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheRoleMenuIdPrefix = "cache:roleMenu:id:"
)

type (
	roleMenuModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleMenuModel)
		Insert(ctx context.Context, in *RoleMenu) (*RoleMenu, error)
		InsertBatch(ctx context.Context, in ...*RoleMenu) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleMenu, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*RoleMenu, err error)
		FindOne(ctx context.Context, id int64) (*RoleMenu, error)
		Update(ctx context.Context, data *RoleMenu) (*RoleMenu, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *RoleMenu) (*RoleMenu, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultRoleMenuModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	RoleMenu struct {
		Id     int64 `gorm:"column:id"`      // 主键id
		RoleId int64 `gorm:"column:role_id"` // 角色id
		MenuId int64 `gorm:"column:menu_id"` // 菜单id
	}
)

func newRoleMenuModel(db *gorm.DB, cache *redis.Client) *defaultRoleMenuModel {
	return &defaultRoleMenuModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`role_menu`",
	}
}

// 切换事务操作
func (m *defaultRoleMenuModel) WithTransaction(tx *gorm.DB) (out RoleMenuModel) {
	return NewRoleMenuModel(tx, m.CacheEngin)
}

// 插入RoleMenu记录
func (m *defaultRoleMenuModel) Insert(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入RoleMenu记录
func (m *defaultRoleMenuModel) InsertBatch(ctx context.Context, in ...*RoleMenu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新RoleMenu记录
func (m *defaultRoleMenuModel) Update(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新RoleMenu记录
func (m *defaultRoleMenuModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新RoleMenu记录
func (m *defaultRoleMenuModel) Save(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除RoleMenu记录
func (m *defaultRoleMenuModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&RoleMenu{})
	return query.RowsAffected, query.Error
}

// 删除RoleMenu记录
func (m *defaultRoleMenuModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&RoleMenu{})
	return result.RowsAffected, result.Error
}

// 查询RoleMenu记录
func (m *defaultRoleMenuModel) First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(RoleMenu)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询RoleMenu总数
func (m *defaultRoleMenuModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&RoleMenu{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询RoleMenu列表
func (m *defaultRoleMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询RoleMenu记录
func (m *defaultRoleMenuModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*RoleMenu, err error) {
	// 创建db
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有搜索条件
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sorts)
	}

	// 如果有分页参数
	if limit > 0 || offset > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询RoleMenu记录
func (m *defaultRoleMenuModel) FindOne(ctx context.Context, id int64) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultRoleMenuModel) TableName() string {
	return m.table
}
