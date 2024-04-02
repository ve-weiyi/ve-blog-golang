package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ FriendLinkModel = (*defaultFriendLinkModel)(nil)

type (
	// 接口定义
	FriendLinkModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out FriendLinkModel)
		// 插入
		Insert(ctx context.Context, in *FriendLink) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*FriendLink) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *FriendLink) (rows int64, err error)
		Save(ctx context.Context, in *FriendLink) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *FriendLink, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *FriendLink, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*FriendLink, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*FriendLink, err error)
		// add extra method in here
	}

	// 表字段定义
	FriendLink struct {
		Id          int64     `json:"id" gorm:"column:id" `                     // id
		LinkName    string    `json:"link_name" gorm:"column:link_name" `       // 链接名
		LinkAvatar  string    `json:"link_avatar" gorm:"column:link_avatar" `   // 链接头像
		LinkAddress string    `json:"link_address" gorm:"column:link_address" ` // 链接地址
		LinkIntro   string    `json:"link_intro" gorm:"column:link_intro" `     // 链接介绍
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at" `     // 创建时间
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at" `     // 更新时间
	}

	// 接口实现
	defaultFriendLinkModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewFriendLinkModel(db *gorm.DB, cache *redis.Client) FriendLinkModel {
	return &defaultFriendLinkModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`friend_link`",
	}
}

// 切换事务操作
func (m *defaultFriendLinkModel) WithTransaction(tx *gorm.DB) (out FriendLinkModel) {
	return NewFriendLinkModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultFriendLinkModel) Insert(ctx context.Context, in *FriendLink) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultFriendLinkModel) InsertBatch(ctx context.Context, in ...*FriendLink) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultFriendLinkModel) Update(ctx context.Context, in *FriendLink) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultFriendLinkModel) Save(ctx context.Context, in *FriendLink) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultFriendLinkModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&FriendLink{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultFriendLinkModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&FriendLink{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultFriendLinkModel) FindOne(ctx context.Context, id int64) (out *FriendLink, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultFriendLinkModel) First(ctx context.Context, conditions string, args ...interface{}) (out *FriendLink, err error) {
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
func (m *defaultFriendLinkModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&FriendLink{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultFriendLinkModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*FriendLink, err error) {
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
func (m *defaultFriendLinkModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*FriendLink, err error) {
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
