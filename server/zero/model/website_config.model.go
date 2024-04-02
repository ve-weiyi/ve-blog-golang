package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameWebsiteConfig = "website_config"

type (
	// 接口定义
	IWebsiteConfigModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out IWebsiteConfigModel)
		// 增删改查
		Create(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error)
		Update(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *WebsiteConfig, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*WebsiteConfig) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*WebsiteConfig, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*WebsiteConfig, err error)
	}

	// 接口实现
	defaultWebsiteConfigModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	WebsiteConfig struct {
		ID        int64  `json:"id"`         // id
		Key       string `json:"key"`        // 关键词
		Config    string `json:"config"`     // 配置信息
		CreatedAt int64  `json:"created_at"` // 创建时间
		UpdatedAt int64  `json:"updated_at"` // 更新时间
	}
)

func NewWebsiteConfigModel(db *gorm.DB, cache *redis.Client) IWebsiteConfigModel {
	return &defaultWebsiteConfigModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameWebsiteConfig,
	}
}

// 切换事务操作
func (s *defaultWebsiteConfigModel) WithTransaction(tx *gorm.DB) (out IWebsiteConfigModel) {
	return NewWebsiteConfigModel(tx, s.CacheEngin)
}

// 创建WebsiteConfig记录
func (s *defaultWebsiteConfigModel) Create(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新WebsiteConfig记录
func (s *defaultWebsiteConfigModel) Update(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除WebsiteConfig记录
func (s *defaultWebsiteConfigModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&WebsiteConfig{})
	return query.RowsAffected, query.Error
}

// 查询WebsiteConfig记录
func (s *defaultWebsiteConfigModel) First(ctx context.Context, conditions string, args ...interface{}) (out *WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(WebsiteConfig)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询WebsiteConfig记录
func (s *defaultWebsiteConfigModel) BatchCreate(ctx context.Context, in ...*WebsiteConfig) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询WebsiteConfig记录
func (s *defaultWebsiteConfigModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&WebsiteConfig{})
	return query.RowsAffected, query.Error
}

// 查询WebsiteConfig总数
func (s *defaultWebsiteConfigModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&WebsiteConfig{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询WebsiteConfig列表
func (s *defaultWebsiteConfigModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*WebsiteConfig, err error) {
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

// 分页查询WebsiteConfig记录
func (s *defaultWebsiteConfigModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*WebsiteConfig, err error) {
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
