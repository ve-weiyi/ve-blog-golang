// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheWebsiteConfigIdPrefix  = "cache:websiteConfig:id:"
	cacheWebsiteConfigKeyPrefix = "cache:websiteConfig:key:"
)

type (
	websiteConfigModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out WebsiteConfigModel)
		Insert(ctx context.Context, in *WebsiteConfig) (*WebsiteConfig, error)
		InsertBatch(ctx context.Context, in ...*WebsiteConfig) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *WebsiteConfig, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*WebsiteConfig, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*WebsiteConfig, err error)
		FindOne(ctx context.Context, id int64) (*WebsiteConfig, error)
		FindOneByKey(ctx context.Context, key string) (out *WebsiteConfig, err error)
		Update(ctx context.Context, data *WebsiteConfig) (*WebsiteConfig, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *WebsiteConfig) (*WebsiteConfig, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultWebsiteConfigModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	WebsiteConfig struct {
		Id        int64     `gorm:"column:id"`         // id
		Key       string    `gorm:"column:key"`        // 关键词
		Config    string    `gorm:"column:config"`     // 配置信息
		CreatedAt time.Time `gorm:"column:created_at"` // 创建时间
		UpdatedAt time.Time `gorm:"column:updated_at"` // 更新时间
	}
)

func newWebsiteConfigModel(db *gorm.DB, cache *redis.Client) *defaultWebsiteConfigModel {
	return &defaultWebsiteConfigModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`website_config`",
	}
}

// 切换事务操作
func (m *defaultWebsiteConfigModel) WithTransaction(tx *gorm.DB) (out WebsiteConfigModel) {
	return NewWebsiteConfigModel(tx, m.CacheEngin)
}

// 插入WebsiteConfig记录
func (m *defaultWebsiteConfigModel) Insert(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入WebsiteConfig记录
func (m *defaultWebsiteConfigModel) InsertBatch(ctx context.Context, in ...*WebsiteConfig) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新WebsiteConfig记录
func (m *defaultWebsiteConfigModel) Update(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新WebsiteConfig记录
func (m *defaultWebsiteConfigModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新WebsiteConfig记录
func (m *defaultWebsiteConfigModel) Save(ctx context.Context, in *WebsiteConfig) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除WebsiteConfig记录
func (m *defaultWebsiteConfigModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&WebsiteConfig{})
	return query.RowsAffected, query.Error
}

// 删除WebsiteConfig记录
func (m *defaultWebsiteConfigModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&WebsiteConfig{})
	return result.RowsAffected, result.Error
}

// 查询WebsiteConfig记录
func (m *defaultWebsiteConfigModel) First(ctx context.Context, conditions string, args ...interface{}) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询WebsiteConfig总数
func (m *defaultWebsiteConfigModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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
func (m *defaultWebsiteConfigModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*WebsiteConfig, err error) {
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

// 分页查询WebsiteConfig记录
func (m *defaultWebsiteConfigModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*WebsiteConfig, err error) {
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

// 查询WebsiteConfig记录
func (m *defaultWebsiteConfigModel) FindOne(ctx context.Context, id int64) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultWebsiteConfigModel) FindOneByKey(ctx context.Context, key string) (out *WebsiteConfig, err error) {
	db := m.DbEngin.WithContext(ctx)

	err = db.Where("`key` = ?", key).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (m *defaultWebsiteConfigModel) TableName() string {
	return m.table
}
