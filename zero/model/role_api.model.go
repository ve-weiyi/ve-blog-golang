package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameRoleApi = "role_api"

type (
	// 接口定义
	RoleApiModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleApiModel)
		// 增删改查
		Create(ctx context.Context, in *RoleApi) (out *RoleApi, err error)
		Update(ctx context.Context, in *RoleApi) (out *RoleApi, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *RoleApi, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*RoleApi) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*RoleApi, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*RoleApi, err error)
	}

	// 接口实现
	defaultRoleApiModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	RoleApi struct {
		Id     int64 `json:"id"`      // 主键id
		RoleId int64 `json:"role_id"` // 角色id
		ApiId  int64 `json:"api_id"`  // 接口id
	}
)

func NewRoleApiModel(db *gorm.DB, cache *redis.Client) RoleApiModel {
	return &defaultRoleApiModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameRoleApi,
	}
}

// 切换事务操作
func (s *defaultRoleApiModel) WithTransaction(tx *gorm.DB) (out RoleApiModel) {
	return NewRoleApiModel(tx, s.CacheEngin)
}

// 创建RoleApi记录
func (s *defaultRoleApiModel) Create(ctx context.Context, in *RoleApi) (out *RoleApi, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新RoleApi记录
func (s *defaultRoleApiModel) Update(ctx context.Context, in *RoleApi) (out *RoleApi, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除RoleApi记录
func (s *defaultRoleApiModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&RoleApi{})
	return query.RowsAffected, query.Error
}

// 查询RoleApi记录
func (s *defaultRoleApiModel) First(ctx context.Context, conditions string, args ...interface{}) (out *RoleApi, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(RoleApi)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询RoleApi记录
func (s *defaultRoleApiModel) BatchCreate(ctx context.Context, in ...*RoleApi) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询RoleApi记录
func (s *defaultRoleApiModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&RoleApi{})
	return query.RowsAffected, query.Error
}

// 查询RoleApi总数
func (s *defaultRoleApiModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&RoleApi{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询RoleApi列表
func (s *defaultRoleApiModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleApi, err error) {
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

// 分页查询RoleApi记录
func (s *defaultRoleApiModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*RoleApi, err error) {
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
