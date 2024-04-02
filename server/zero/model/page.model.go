package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNamePage = "page"

type (
	// 接口定义
	IPageModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out IPageModel)
		// 增删改查
		Create(ctx context.Context, in *Page) (out *Page, err error)
		Update(ctx context.Context, in *Page) (out *Page, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Page, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Page) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Page, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Page, err error)
	}

	// 接口实现
	defaultPageModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Page struct {
		ID        int64  `json:"id"`         // 页面id
		PageName  string `json:"page_name"`  // 页面名
		PageLabel string `json:"page_label"` // 页面标签
		PageCover string `json:"page_cover"` // 页面封面
		CreatedAt int64  `json:"created_at"` // 创建时间
		UpdatedAt int64  `json:"updated_at"` // 更新时间
	}
)

func NewPageModel(db *gorm.DB, cache *redis.Client) IPageModel {
	return &defaultPageModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNamePage,
	}
}

// 切换事务操作
func (s *defaultPageModel) WithTransaction(tx *gorm.DB) (out IPageModel) {
	return NewPageModel(tx, s.CacheEngin)
}

// 创建Page记录
func (s *defaultPageModel) Create(ctx context.Context, in *Page) (out *Page, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Page记录
func (s *defaultPageModel) Update(ctx context.Context, in *Page) (out *Page, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Page记录
func (s *defaultPageModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Page{})
	return query.RowsAffected, query.Error
}

// 查询Page记录
func (s *defaultPageModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Page, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Page)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Page记录
func (s *defaultPageModel) BatchCreate(ctx context.Context, in ...*Page) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Page记录
func (s *defaultPageModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Page{})
	return query.RowsAffected, query.Error
}

// 查询Page总数
func (s *defaultPageModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Page{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Page列表
func (s *defaultPageModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Page, err error) {
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

// 分页查询Page记录
func (s *defaultPageModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Page, err error) {
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
	if page > 0 && size > 0 {
		limit := size
		offset := (page - 1) * limit
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
