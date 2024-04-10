package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameTalk = "talk"

type (
	// 接口定义
	TalkModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out TalkModel)
		// 增删改查
		Create(ctx context.Context, in *Talk) (out *Talk, err error)
		Update(ctx context.Context, in *Talk) (out *Talk, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Talk, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Talk) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Talk, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Talk, err error)
	}

	// 接口实现
	defaultTalkModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Talk struct {
		Id        int64     `json:"id"`         // 说说id
		UserId    int64     `json:"user_id"`    // 用户id
		Content   string    `json:"content"`    // 说说内容
		Images    string    `json:"images"`     // 图片
		IsTop     int64     `json:"is_top"`     // 是否置顶
		Status    int64     `json:"status"`     // 状态 1.公开 2.私密
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewTalkModel(db *gorm.DB, cache *redis.Client) TalkModel {
	return &defaultTalkModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameTalk,
	}
}

// 切换事务操作
func (s *defaultTalkModel) WithTransaction(tx *gorm.DB) (out TalkModel) {
	return NewTalkModel(tx, s.CacheEngin)
}

// 创建Talk记录
func (s *defaultTalkModel) Create(ctx context.Context, in *Talk) (out *Talk, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Talk记录
func (s *defaultTalkModel) Update(ctx context.Context, in *Talk) (out *Talk, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Talk记录
func (s *defaultTalkModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Talk{})
	return query.RowsAffected, query.Error
}

// 查询Talk记录
func (s *defaultTalkModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Talk, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Talk)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Talk记录
func (s *defaultTalkModel) BatchCreate(ctx context.Context, in ...*Talk) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Talk记录
func (s *defaultTalkModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Talk{})
	return query.RowsAffected, query.Error
}

// 查询Talk总数
func (s *defaultTalkModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Talk{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Talk列表
func (s *defaultTalkModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Talk, err error) {
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

// 分页查询Talk记录
func (s *defaultTalkModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Talk, err error) {
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
