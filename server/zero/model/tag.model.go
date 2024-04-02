package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameTag = "tag"

type (
	// 接口定义
	ITagModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ITagModel)
		// 增删改查
		Create(ctx context.Context, in *Tag) (out *Tag, err error)
		Update(ctx context.Context, in *Tag) (out *Tag, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Tag, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Tag) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Tag, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Tag, err error)
	}

	// 接口实现
	defaultTagModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Tag struct {
		ID        int64  `json:"id"`         // id
		TagName   string `json:"tag_name"`   // 标签名
		CreatedAt int64  `json:"created_at"` // 创建时间
		UpdatedAt int64  `json:"updated_at"` // 更新时间
	}
)

func NewTagModel(db *gorm.DB, cache *redis.Client) ITagModel {
	return &defaultTagModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameTag,
	}
}

// 切换事务操作
func (s *defaultTagModel) WithTransaction(tx *gorm.DB) (out ITagModel) {
	return NewTagModel(tx, s.CacheEngin)
}

// 创建Tag记录
func (s *defaultTagModel) Create(ctx context.Context, in *Tag) (out *Tag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Tag记录
func (s *defaultTagModel) Update(ctx context.Context, in *Tag) (out *Tag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Tag记录
func (s *defaultTagModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Tag{})
	return query.RowsAffected, query.Error
}

// 查询Tag记录
func (s *defaultTagModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Tag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询Tag记录
func (s *defaultTagModel) BatchCreate(ctx context.Context, in ...*Tag) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Tag记录
func (s *defaultTagModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Tag{})
	return query.RowsAffected, query.Error
}

// 查询Tag总数
func (s *defaultTagModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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
func (s *defaultTagModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Tag, err error) {
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

// 分页查询Tag记录
func (s *defaultTagModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Tag, err error) {
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
