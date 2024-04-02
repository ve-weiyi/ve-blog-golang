package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameApi = "api"

type (
	// 接口定义
	ApiModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ApiModel)
		// 增删改查
		Create(ctx context.Context, in *Api) (out *Api, err error)
		Update(ctx context.Context, in *Api) (out *Api, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Api, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Api) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Api, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Api, err error)
	}

	// 接口实现
	defaultApiModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Api struct {
		Id        int64     `json:"id"`         // 主键id
		ParentId  int64     `json:"parent_id"`  // 分组id
		Name      string    `json:"name"`       // api名称
		Path      string    `json:"path"`       // api路径
		Method    string    `json:"method"`     // api请求方法
		Traceable int64     `json:"traceable"`  // 是否追溯操作记录 0需要，1是
		Status    int64     `json:"status"`     // 状态 1开，2关
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewApiModel(db *gorm.DB, cache *redis.Client) ApiModel {
	return &defaultApiModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameApi,
	}
}

// 切换事务操作
func (s *defaultApiModel) WithTransaction(tx *gorm.DB) (out ApiModel) {
	return NewApiModel(tx, s.CacheEngin)
}

// 创建Api记录
func (s *defaultApiModel) Create(ctx context.Context, in *Api) (out *Api, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Api记录
func (s *defaultApiModel) Update(ctx context.Context, in *Api) (out *Api, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Api记录
func (s *defaultApiModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Api{})
	return query.RowsAffected, query.Error
}

// 查询Api记录
func (s *defaultApiModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Api, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Api)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Api记录
func (s *defaultApiModel) BatchCreate(ctx context.Context, in ...*Api) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Api记录
func (s *defaultApiModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Api{})
	return query.RowsAffected, query.Error
}

// 查询Api总数
func (s *defaultApiModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Api{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Api列表
func (s *defaultApiModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Api, err error) {
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

// 分页查询Api记录
func (s *defaultApiModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Api, err error) {
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
