// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheCategoryIdPrefix = "cache:category:id:"
)

type (
	categoryModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out CategoryModel)
		Insert(ctx context.Context, in *Category) (*Category, error)
		InsertBatch(ctx context.Context, in ...*Category) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Category, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Category, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*Category, err error)
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) (*Category, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *Category) (*Category, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultCategoryModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	Category struct {
		Id           int64     `gorm:"column:id"`            // id
		CategoryName string    `gorm:"column:category_name"` // 分类名
		CreatedAt    time.Time `gorm:"column:created_at"`    // 创建时间
		UpdatedAt    time.Time `gorm:"column:updated_at"`    // 更新时间
	}
)

func newCategoryModel(db *gorm.DB, cache *redis.Client) *defaultCategoryModel {
	return &defaultCategoryModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`category`",
	}
}

// 切换事务操作
func (m *defaultCategoryModel) WithTransaction(tx *gorm.DB) (out CategoryModel) {
	return NewCategoryModel(tx, m.CacheEngin)
}

// 插入Category记录
func (m *defaultCategoryModel) Insert(ctx context.Context, in *Category) (out *Category, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入Category记录
func (m *defaultCategoryModel) InsertBatch(ctx context.Context, in ...*Category) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新Category记录
func (m *defaultCategoryModel) Update(ctx context.Context, in *Category) (out *Category, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Category记录
func (m *defaultCategoryModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新Category记录
func (m *defaultCategoryModel) Save(ctx context.Context, in *Category) (out *Category, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Category记录
func (m *defaultCategoryModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&Category{})
	return query.RowsAffected, query.Error
}

// 删除Category记录
func (m *defaultCategoryModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Category{})
	return result.RowsAffected, result.Error
}

// 查询Category记录
func (m *defaultCategoryModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Category, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Category)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Category总数
func (m *defaultCategoryModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Category{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Category列表
func (m *defaultCategoryModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Category, err error) {
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

// 分页查询Category记录
func (m *defaultCategoryModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Category, err error) {
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

// 查询Category记录
func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (out *Category, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultCategoryModel) TableName() string {
	return m.table
}
