package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameRoleMenu = "role_menu"

type (
	// 接口定义
	RoleMenuModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleMenuModel)
		// 增删改查
		Create(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error)
		Update(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*RoleMenu) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*RoleMenu, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*RoleMenu, err error)
	}

	// 接口实现
	defaultRoleMenuModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	RoleMenu struct {
		Id     int64 `json:"id"`      // 主键id
		RoleId int64 `json:"role_id"` // 角色id
		MenuId int64 `json:"menu_id"` // 菜单id
	}
)

func NewRoleMenuModel(db *gorm.DB, cache *redis.Client) RoleMenuModel {
	return &defaultRoleMenuModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameRoleMenu,
	}
}

// 切换事务操作
func (s *defaultRoleMenuModel) WithTransaction(tx *gorm.DB) (out RoleMenuModel) {
	return NewRoleMenuModel(tx, s.CacheEngin)
}

// 创建RoleMenu记录
func (s *defaultRoleMenuModel) Create(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新RoleMenu记录
func (s *defaultRoleMenuModel) Update(ctx context.Context, in *RoleMenu) (out *RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除RoleMenu记录
func (s *defaultRoleMenuModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&RoleMenu{})
	return query.RowsAffected, query.Error
}

// 查询RoleMenu记录
func (s *defaultRoleMenuModel) First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询RoleMenu记录
func (s *defaultRoleMenuModel) BatchCreate(ctx context.Context, in ...*RoleMenu) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询RoleMenu记录
func (s *defaultRoleMenuModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&RoleMenu{})
	return query.RowsAffected, query.Error
}

// 查询RoleMenu总数
func (s *defaultRoleMenuModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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
func (s *defaultRoleMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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
func (s *defaultRoleMenuModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*RoleMenu, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有搜索条件
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sorts)
	}

	// 如果有分页参数
	if limit > 0 && offset > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
