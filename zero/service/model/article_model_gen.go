package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ArticleModel = (*defaultArticleModel)(nil)

type (
	// 接口定义
	ArticleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ArticleModel)
		// 插入
		Insert(ctx context.Context, in *Article) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Article) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *Article) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *Article) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Article, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Article, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Article, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Article, err error)
		// add extra method in here
	}

	// 表字段定义
	Article struct {
		Id             int64     `json:"id" gorm:"column:id" `                           // id
		UserId         int64     `json:"user_id" gorm:"column:user_id" `                 // 作者
		CategoryId     int64     `json:"category_id" gorm:"column:category_id" `         // 文章分类
		ArticleCover   string    `json:"article_cover" gorm:"column:article_cover" `     // 文章缩略图
		ArticleTitle   string    `json:"article_title" gorm:"column:article_title" `     // 标题
		ArticleContent string    `json:"article_content" gorm:"column:article_content" ` // 内容
		ArticleType    int64     `json:"article_type" gorm:"column:article_type" `       // 文章类型 1原创 2转载 3翻译
		OriginalUrl    string    `json:"original_url" gorm:"column:original_url" `       // 原文链接
		IsTop          int64     `json:"is_top" gorm:"column:is_top" `                   // 是否置顶 0否 1是
		IsDelete       int64     `json:"is_delete" gorm:"column:is_delete" `             // 是否删除  0否 1是
		Status         int64     `json:"status" gorm:"column:status" `                   // 状态值 1公开 2私密 3评论可见
		LikeCount      int64     `json:"like_count" gorm:"column:like_count" `           // 点赞数
		CreatedAt      time.Time `json:"created_at" gorm:"column:created_at" `           // 发表时间
		UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at" `           // 更新时间
	}

	// 接口实现
	defaultArticleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewArticleModel(db *gorm.DB, cache *redis.Client) ArticleModel {
	return &defaultArticleModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`article`",
	}
}

// 切换事务操作
func (m *defaultArticleModel) WithTransaction(tx *gorm.DB) (out ArticleModel) {
	return NewArticleModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultArticleModel) Insert(ctx context.Context, in *Article) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultArticleModel) InsertBatch(ctx context.Context, in ...*Article) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultArticleModel) Update(ctx context.Context, in *Article) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultArticleModel) UpdateNotEmpty(ctx context.Context, in *Article) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultArticleModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Article{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultArticleModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Article{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (out *Article, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultArticleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Article, err error) {
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
func (m *defaultArticleModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询列表
func (m *defaultArticleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Article, err error) {
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
func (m *defaultArticleModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Article, err error) {
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
