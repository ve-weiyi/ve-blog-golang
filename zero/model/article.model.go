package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameArticle = "article"

type (
	// 接口定义
	ArticleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ArticleModel)
		// 增删改查
		Create(ctx context.Context, in *Article) (out *Article, err error)
		Update(ctx context.Context, in *Article) (out *Article, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Article, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Article) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Article, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Article, err error)
	}

	// 接口实现
	defaultArticleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Article struct {
		Id             int64     `json:"id"`              // id
		UserId         int64     `json:"user_id"`         // 作者
		CategoryId     int64     `json:"category_id"`     // 文章分类
		ArticleCover   string    `json:"article_cover"`   // 文章缩略图
		ArticleTitle   string    `json:"article_title"`   // 标题
		ArticleContent string    `json:"article_content"` // 内容
		Type           int64     `json:"type"`            // 文章类型 1原创 2转载 3翻译
		OriginalUrl    string    `json:"original_url"`    // 原文链接
		IsTop          int64     `json:"is_top"`          // 是否置顶 0否 1是
		IsDelete       int64     `json:"is_delete"`       // 是否删除  0否 1是
		Status         int64     `json:"status"`          // 状态值 1公开 2私密 3评论可见
		CreatedAt      time.Time `json:"created_at"`      // 发表时间
		UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	}
)

func NewArticleModel(db *gorm.DB, cache *redis.Client) ArticleModel {
	return &defaultArticleModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameArticle,
	}
}

// 切换事务操作
func (s *defaultArticleModel) WithTransaction(tx *gorm.DB) (out ArticleModel) {
	return NewArticleModel(tx, s.CacheEngin)
}

// 创建Article记录
func (s *defaultArticleModel) Create(ctx context.Context, in *Article) (out *Article, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Article记录
func (s *defaultArticleModel) Update(ctx context.Context, in *Article) (out *Article, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Article记录
func (s *defaultArticleModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Article{})
	return query.RowsAffected, query.Error
}

// 查询Article记录
func (s *defaultArticleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Article, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Article)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Article记录
func (s *defaultArticleModel) BatchCreate(ctx context.Context, in ...*Article) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Article记录
func (s *defaultArticleModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Article{})
	return query.RowsAffected, query.Error
}

// 查询Article总数
func (s *defaultArticleModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Article列表
func (s *defaultArticleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Article, err error) {
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

// 分页查询Article记录
func (s *defaultArticleModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Article, err error) {
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
