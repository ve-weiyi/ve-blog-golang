// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheUniqueViewIdPrefix = "cache:uniqueView:id:"
)

type (
	uniqueViewModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UniqueViewModel)
		Insert(ctx context.Context, in *UniqueView) (*UniqueView, error)
		InsertBatch(ctx context.Context, in ...*UniqueView) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UniqueView, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UniqueView, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*UniqueView, err error)
		FindOne(ctx context.Context, id int64) (*UniqueView, error)
		Update(ctx context.Context, data *UniqueView) (*UniqueView, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *UniqueView) (*UniqueView, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultUniqueViewModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	UniqueView struct {
		Id         int64     `gorm:"column:id"`          // id
		ViewsCount int64     `gorm:"column:views_count"` // 访问量
		CreatedAt  time.Time `gorm:"column:created_at"`  // 创建时间
		UpdatedAt  time.Time `gorm:"column:updated_at"`  // 更新时间
	}
)

func newUniqueViewModel(db *gorm.DB, cache *redis.Client) *defaultUniqueViewModel {
	return &defaultUniqueViewModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`unique_view`",
	}
}

// 切换事务操作
func (m *defaultUniqueViewModel) WithTransaction(tx *gorm.DB) (out UniqueViewModel) {
	return NewUniqueViewModel(tx, m.CacheEngin)
}

// 插入UniqueView记录
func (m *defaultUniqueViewModel) Insert(ctx context.Context, in *UniqueView) (out *UniqueView, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入UniqueView记录
func (m *defaultUniqueViewModel) InsertBatch(ctx context.Context, in ...*UniqueView) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新UniqueView记录
func (m *defaultUniqueViewModel) Update(ctx context.Context, in *UniqueView) (out *UniqueView, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UniqueView记录
func (m *defaultUniqueViewModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新UniqueView记录
func (m *defaultUniqueViewModel) Save(ctx context.Context, in *UniqueView) (out *UniqueView, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UniqueView记录
func (m *defaultUniqueViewModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&UniqueView{})
	return query.RowsAffected, query.Error
}

// 删除UniqueView记录
func (m *defaultUniqueViewModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&UniqueView{})
	return result.RowsAffected, result.Error
}

// 查询UniqueView记录
func (m *defaultUniqueViewModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UniqueView, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询UniqueView总数
func (m *defaultUniqueViewModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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
func (m *defaultUniqueViewModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UniqueView, err error) {
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

// 分页查询UniqueView记录
func (m *defaultUniqueViewModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UniqueView, err error) {
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

// 查询UniqueView记录
func (m *defaultUniqueViewModel) FindOne(ctx context.Context, id int64) (out *UniqueView, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultUniqueViewModel) TableName() string {
	return m.table
}
