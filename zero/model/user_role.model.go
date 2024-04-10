package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUserRole = "user_role"

type (
	// 接口定义
	UserRoleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserRoleModel)
		// 增删改查
		Create(ctx context.Context, in *UserRole) (out *UserRole, err error)
		Update(ctx context.Context, in *UserRole) (out *UserRole, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserRole, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UserRole) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserRole, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserRole, err error)
	}

	// 接口实现
	defaultUserRoleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UserRole struct {
		Id     int64 `json:"id"`      // 主键id
		UserId int64 `json:"user_id"` // 用户id
		RoleId int64 `json:"role_id"` // 角色id
	}
)

func NewUserRoleModel(db *gorm.DB, cache *redis.Client) UserRoleModel {
	return &defaultUserRoleModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUserRole,
	}
}

// 切换事务操作
func (s *defaultUserRoleModel) WithTransaction(tx *gorm.DB) (out UserRoleModel) {
	return NewUserRoleModel(tx, s.CacheEngin)
}

// 创建UserRole记录
func (s *defaultUserRoleModel) Create(ctx context.Context, in *UserRole) (out *UserRole, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UserRole记录
func (s *defaultUserRoleModel) Update(ctx context.Context, in *UserRole) (out *UserRole, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UserRole记录
func (s *defaultUserRoleModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UserRole{})
	return query.RowsAffected, query.Error
}

// 查询UserRole记录
func (s *defaultUserRoleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserRole, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UserRole)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UserRole记录
func (s *defaultUserRoleModel) BatchCreate(ctx context.Context, in ...*UserRole) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UserRole记录
func (s *defaultUserRoleModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UserRole{})
	return query.RowsAffected, query.Error
}

// 查询UserRole总数
func (s *defaultUserRoleModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UserRole{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserRole列表
func (s *defaultUserRoleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserRole, err error) {
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

// 分页查询UserRole记录
func (s *defaultUserRoleModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserRole, err error) {
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
