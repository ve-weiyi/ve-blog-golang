package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameCategory = "category"

type (
	// 接口定义
	CategoryModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out CategoryModel)
		// 增删改查
		Create(ctx context.Context, in *Category) (out *Category, err error)
		Update(ctx context.Context, in *Category) (out *Category, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Category, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Category) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Category, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Category, err error)
	}

	// 接口实现
	defaultCategoryModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Category struct {
		Id           int64     `json:"id"`            // id
		CategoryName string    `json:"category_name"` // 分类名
		CreatedAt    time.Time `json:"created_at"`    // 创建时间
		UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
	}
)

func NewCategoryModel(db *gorm.DB, cache *redis.Client) CategoryModel {
	return &defaultCategoryModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameCategory,
	}
}

// 切换事务操作
func (s *defaultCategoryModel) WithTransaction(tx *gorm.DB) (out CategoryModel) {
	return NewCategoryModel(tx, s.CacheEngin)
}

// 创建Category记录
func (s *defaultCategoryModel) Create(ctx context.Context, in *Category) (out *Category, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Category记录
func (s *defaultCategoryModel) Update(ctx context.Context, in *Category) (out *Category, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Category记录
func (s *defaultCategoryModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Category{})
	return query.RowsAffected, query.Error
}

// 查询Category记录
func (s *defaultCategoryModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Category, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询Category记录
func (s *defaultCategoryModel) BatchCreate(ctx context.Context, in ...*Category) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Category记录
func (s *defaultCategoryModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Category{})
	return query.RowsAffected, query.Error
}

// 查询Category总数
func (s *defaultCategoryModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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
func (s *defaultCategoryModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Category, err error) {
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

// 分页查询Category记录
func (s *defaultCategoryModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Category, err error) {
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
