package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUniqueView = "unique_view"

type (
	// 接口定义
	UniqueViewModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UniqueViewModel)
		// 增删改查
		Create(ctx context.Context, in *UniqueView) (out *UniqueView, err error)
		Update(ctx context.Context, in *UniqueView) (out *UniqueView, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UniqueView, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UniqueView) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UniqueView, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UniqueView, err error)
	}

	// 接口实现
	defaultUniqueViewModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UniqueView struct {
		Id         int64     `json:"id"`          // id
		ViewsCount int64     `json:"views_count"` // 访问量
		CreatedAt  time.Time `json:"created_at"`  // 创建时间
		UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
	}
)

func NewUniqueViewModel(db *gorm.DB, cache *redis.Client) UniqueViewModel {
	return &defaultUniqueViewModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUniqueView,
	}
}

// 切换事务操作
func (s *defaultUniqueViewModel) WithTransaction(tx *gorm.DB) (out UniqueViewModel) {
	return NewUniqueViewModel(tx, s.CacheEngin)
}

// 创建UniqueView记录
func (s *defaultUniqueViewModel) Create(ctx context.Context, in *UniqueView) (out *UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UniqueView记录
func (s *defaultUniqueViewModel) Update(ctx context.Context, in *UniqueView) (out *UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UniqueView记录
func (s *defaultUniqueViewModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UniqueView{})
	return query.RowsAffected, query.Error
}

// 查询UniqueView记录
func (s *defaultUniqueViewModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UniqueView)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UniqueView记录
func (s *defaultUniqueViewModel) BatchCreate(ctx context.Context, in ...*UniqueView) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UniqueView记录
func (s *defaultUniqueViewModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UniqueView{})
	return query.RowsAffected, query.Error
}

// 查询UniqueView总数
func (s *defaultUniqueViewModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UniqueView{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UniqueView列表
func (s *defaultUniqueViewModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UniqueView, err error) {
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

// 分页查询UniqueView记录
func (s *defaultUniqueViewModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UniqueView, err error) {
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
