package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameRole = "role"

type (
	// 接口定义
	RoleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleModel)
		// 增删改查
		Create(ctx context.Context, in *Role) (out *Role, err error)
		Update(ctx context.Context, in *Role) (out *Role, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Role, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Role) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Role, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Role, err error)
	}

	// 接口实现
	defaultRoleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Role struct {
		Id          int64     `json:"id"`           // 主键id
		ParentId    int64     `json:"parent_id"`    // 父角色id
		RoleDomain  string    `json:"role_domain"`  // 角色域
		RoleName    string    `json:"role_name"`    // 角色名
		RoleComment string    `json:"role_comment"` // 角色备注
		IsDisable   int64     `json:"is_disable"`   // 是否禁用  0否 1是
		IsDefault   int64     `json:"is_default"`   // 是否默认角色 0否 1是
		CreatedAt   time.Time `json:"created_at"`   // 创建时间
		UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
	}
)

func NewRoleModel(db *gorm.DB, cache *redis.Client) RoleModel {
	return &defaultRoleModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameRole,
	}
}

// 切换事务操作
func (s *defaultRoleModel) WithTransaction(tx *gorm.DB) (out RoleModel) {
	return NewRoleModel(tx, s.CacheEngin)
}

// 创建Role记录
func (s *defaultRoleModel) Create(ctx context.Context, in *Role) (out *Role, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Role记录
func (s *defaultRoleModel) Update(ctx context.Context, in *Role) (out *Role, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Role记录
func (s *defaultRoleModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Role{})
	return query.RowsAffected, query.Error
}

// 查询Role记录
func (s *defaultRoleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Role, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Role)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Role记录
func (s *defaultRoleModel) BatchCreate(ctx context.Context, in ...*Role) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Role记录
func (s *defaultRoleModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Role{})
	return query.RowsAffected, query.Error
}

// 查询Role总数
func (s *defaultRoleModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Role{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Role列表
func (s *defaultRoleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Role, err error) {
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

// 分页查询Role记录
func (s *defaultRoleModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Role, err error) {
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
