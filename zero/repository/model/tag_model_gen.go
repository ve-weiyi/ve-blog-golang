// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheTagIdPrefix = "cache:tag:id:"
)

type (
	tagModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out TagModel)
		Insert(ctx context.Context, in *Tag) (*Tag, error)
		InsertBatch(ctx context.Context, in ...*Tag) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Tag, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Tag, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*Tag, err error)
		FindOne(ctx context.Context, id int64) (*Tag, error)
		Update(ctx context.Context, data *Tag) (*Tag, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *Tag) (*Tag, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultTagModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	Tag struct {
		Id        int64     `gorm:"column:id"`         // id
		TagName   string    `gorm:"column:tag_name"`   // 标签名
		CreatedAt time.Time `gorm:"column:created_at"` // 创建时间
		UpdatedAt time.Time `gorm:"column:updated_at"` // 更新时间
	}
)

func newTagModel(db *gorm.DB, cache *redis.Client) *defaultTagModel {
	return &defaultTagModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`tag`",
	}
}

// 切换事务操作
func (m *defaultTagModel) WithTransaction(tx *gorm.DB) (out TagModel) {
	return NewTagModel(tx, m.CacheEngin)
}

// 插入Tag记录
func (m *defaultTagModel) Insert(ctx context.Context, in *Tag) (out *Tag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入Tag记录
func (m *defaultTagModel) InsertBatch(ctx context.Context, in ...*Tag) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新Tag记录
func (m *defaultTagModel) Update(ctx context.Context, in *Tag) (out *Tag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Tag记录
func (m *defaultTagModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新Tag记录
func (m *defaultTagModel) Save(ctx context.Context, in *Tag) (out *Tag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Tag记录
func (m *defaultTagModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&Tag{})
	return query.RowsAffected, query.Error
}

// 删除Tag记录
func (m *defaultTagModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Tag{})
	return result.RowsAffected, result.Error
}

// 查询Tag记录
func (m *defaultTagModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Tag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Tag)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Tag总数
func (m *defaultTagModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Tag{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Tag列表
func (m *defaultTagModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Tag, err error) {
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

// 分页查询Tag记录
func (m *defaultTagModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Tag, err error) {
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

// 查询Tag记录
func (m *defaultTagModel) FindOne(ctx context.Context, id int64) (out *Tag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultTagModel) TableName() string {
	return m.table
}
