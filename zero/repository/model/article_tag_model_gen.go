// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheArticleTagIdPrefix = "cache:articleTag:id:"
)

type (
	articleTagModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ArticleTagModel)
		Insert(ctx context.Context, in *ArticleTag) (*ArticleTag, error)
		InsertBatch(ctx context.Context, in ...*ArticleTag) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ArticleTag, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ArticleTag, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*ArticleTag, err error)
		FindOne(ctx context.Context, id int64) (*ArticleTag, error)
		Update(ctx context.Context, data *ArticleTag) (*ArticleTag, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *ArticleTag) (*ArticleTag, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultArticleTagModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	ArticleTag struct {
		Id        int64 `gorm:"column:id"`         // id
		ArticleId int64 `gorm:"column:article_id"` // 文章id
		TagId     int64 `gorm:"column:tag_id"`     // 标签id
	}
)

func newArticleTagModel(db *gorm.DB, cache *redis.Client) *defaultArticleTagModel {
	return &defaultArticleTagModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`article_tag`",
	}
}

// 切换事务操作
func (m *defaultArticleTagModel) WithTransaction(tx *gorm.DB) (out ArticleTagModel) {
	return NewArticleTagModel(tx, m.CacheEngin)
}

// 插入ArticleTag记录
func (m *defaultArticleTagModel) Insert(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入ArticleTag记录
func (m *defaultArticleTagModel) InsertBatch(ctx context.Context, in ...*ArticleTag) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新ArticleTag记录
func (m *defaultArticleTagModel) Update(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新ArticleTag记录
func (m *defaultArticleTagModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新ArticleTag记录
func (m *defaultArticleTagModel) Save(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除ArticleTag记录
func (m *defaultArticleTagModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&ArticleTag{})
	return query.RowsAffected, query.Error
}

// 删除ArticleTag记录
func (m *defaultArticleTagModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&ArticleTag{})
	return result.RowsAffected, result.Error
}

// 查询ArticleTag记录
func (m *defaultArticleTagModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ArticleTag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(ArticleTag)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询ArticleTag总数
func (m *defaultArticleTagModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&ArticleTag{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询ArticleTag列表
func (m *defaultArticleTagModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ArticleTag, err error) {
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

// 分页查询ArticleTag记录
func (m *defaultArticleTagModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ArticleTag, err error) {
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

// 查询ArticleTag记录
func (m *defaultArticleTagModel) FindOne(ctx context.Context, id int64) (out *ArticleTag, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultArticleTagModel) TableName() string {
	return m.table
}
