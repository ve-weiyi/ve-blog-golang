package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameFriendLink = "friend_link"

type (
	// 接口定义
	FriendLinkModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out FriendLinkModel)
		// 增删改查
		Create(ctx context.Context, in *FriendLink) (out *FriendLink, err error)
		Update(ctx context.Context, in *FriendLink) (out *FriendLink, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *FriendLink, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*FriendLink) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*FriendLink, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*FriendLink, err error)
	}

	// 接口实现
	defaultFriendLinkModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	FriendLink struct {
		Id          int64     `json:"id"`           // id
		LinkName    string    `json:"link_name"`    // 链接名
		LinkAvatar  string    `json:"link_avatar"`  // 链接头像
		LinkAddress string    `json:"link_address"` // 链接地址
		LinkIntro   string    `json:"link_intro"`   // 链接介绍
		CreatedAt   time.Time `json:"created_at"`   // 创建时间
		UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
	}
)

func NewFriendLinkModel(db *gorm.DB, cache *redis.Client) FriendLinkModel {
	return &defaultFriendLinkModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameFriendLink,
	}
}

// 切换事务操作
func (s *defaultFriendLinkModel) WithTransaction(tx *gorm.DB) (out FriendLinkModel) {
	return NewFriendLinkModel(tx, s.CacheEngin)
}

// 创建FriendLink记录
func (s *defaultFriendLinkModel) Create(ctx context.Context, in *FriendLink) (out *FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新FriendLink记录
func (s *defaultFriendLinkModel) Update(ctx context.Context, in *FriendLink) (out *FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除FriendLink记录
func (s *defaultFriendLinkModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&FriendLink{})
	return query.RowsAffected, query.Error
}

// 查询FriendLink记录
func (s *defaultFriendLinkModel) First(ctx context.Context, conditions string, args ...interface{}) (out *FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(FriendLink)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询FriendLink记录
func (s *defaultFriendLinkModel) BatchCreate(ctx context.Context, in ...*FriendLink) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询FriendLink记录
func (s *defaultFriendLinkModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&FriendLink{})
	return query.RowsAffected, query.Error
}

// 查询FriendLink总数
func (s *defaultFriendLinkModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询FriendLink列表
func (s *defaultFriendLinkModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*FriendLink, err error) {
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

// 分页查询FriendLink记录
func (s *defaultFriendLinkModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*FriendLink, err error) {
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
