package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ BannerModel = (*defaultBannerModel)(nil)

type (
	// 接口定义
	BannerModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out BannerModel)
		// 插入
		Insert(ctx context.Context, in *Banner) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Banner) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *Banner) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *Banner) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Banner, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Banner, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Banner, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Banner, err error)
		// add extra method in here
	}

	// 表字段定义
	Banner struct {
		Id          int64     `json:"id" gorm:"column:id" `                     // 页面id
		BannerName  string    `json:"banner_name" gorm:"column:banner_name" `   // 页面名
		BannerLabel string    `json:"banner_label" gorm:"column:banner_label" ` // 页面标签
		BannerCover string    `json:"banner_cover" gorm:"column:banner_cover" ` // 页面封面
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at" `     // 创建时间
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at" `     // 更新时间
	}

	// 接口实现
	defaultBannerModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewBannerModel(db *gorm.DB, cache *redis.Client) BannerModel {
	return &defaultBannerModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`banner`",
	}
}

// 切换事务操作
func (m *defaultBannerModel) WithTransaction(tx *gorm.DB) (out BannerModel) {
	return NewBannerModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultBannerModel) Insert(ctx context.Context, in *Banner) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultBannerModel) InsertBatch(ctx context.Context, in ...*Banner) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultBannerModel) Update(ctx context.Context, in *Banner) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultBannerModel) UpdateNotEmpty(ctx context.Context, in *Banner) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultBannerModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Banner{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultBannerModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Banner{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultBannerModel) FindOne(ctx context.Context, id int64) (out *Banner, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultBannerModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Banner, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询总数
func (m *defaultBannerModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Banner{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultBannerModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Banner, err error) {
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

// 分页查询记录
func (m *defaultBannerModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Banner, err error) {
	// 插入db
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

// add extra method in here
