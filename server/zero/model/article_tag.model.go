package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameArticleTag = "article_tag"

type (
	// 接口定义
	IArticleTagModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out IArticleTagModel)
		// 增删改查
		Create(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error)
		Update(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ArticleTag, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*ArticleTag) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*ArticleTag, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ArticleTag, err error)
	}

	// 接口实现
	defaultArticleTagModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	ArticleTag struct {
		ID        int64 `json:"id"`         // id
		ArticleID int64 `json:"article_id"` // 文章id
		TagID     int64 `json:"tag_id"`     // 标签id
	}
)

func NewArticleTagModel(db *gorm.DB, cache *redis.Client) IArticleTagModel {
	return &defaultArticleTagModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameArticleTag,
	}
}

// 切换事务操作
func (s *defaultArticleTagModel) WithTransaction(tx *gorm.DB) (out IArticleTagModel) {
	return NewArticleTagModel(tx, s.CacheEngin)
}

// 创建ArticleTag记录
func (s *defaultArticleTagModel) Create(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新ArticleTag记录
func (s *defaultArticleTagModel) Update(ctx context.Context, in *ArticleTag) (out *ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除ArticleTag记录
func (s *defaultArticleTagModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&ArticleTag{})
	return query.RowsAffected, query.Error
}

// 查询ArticleTag记录
func (s *defaultArticleTagModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询ArticleTag记录
func (s *defaultArticleTagModel) BatchCreate(ctx context.Context, in ...*ArticleTag) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询ArticleTag记录
func (s *defaultArticleTagModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&ArticleTag{})
	return query.RowsAffected, query.Error
}

// 查询ArticleTag总数
func (s *defaultArticleTagModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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
func (s *defaultArticleTagModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ArticleTag, err error) {
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

// 分页查询ArticleTag记录
func (s *defaultArticleTagModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*ArticleTag, err error) {
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
