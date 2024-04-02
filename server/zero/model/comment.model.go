package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameComment = "comment"

type (
	// 接口定义
	ICommentModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ICommentModel)
		// 增删改查
		Create(ctx context.Context, in *Comment) (out *Comment, err error)
		Update(ctx context.Context, in *Comment) (out *Comment, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Comment) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Comment, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Comment, err error)
	}

	// 接口实现
	defaultCommentModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Comment struct {
		ID             int64  `json:"id"`              // 主键
		UserID         int64  `json:"user_id"`         // 评论用户Id
		TopicID        int64  `json:"topic_id"`        // 评论主题id
		CommentContent string `json:"comment_content"` // 评论内容
		ReplyUserID    int64  `json:"reply_user_id"`   // 回复用户id
		ParentID       int64  `json:"parent_id"`       // 父评论id
		Type           int64  `json:"type"`            // 评论类型 1.文章 2.友链 3.说说
		IsDelete       int64  `json:"is_delete"`       // 是否删除  0否 1是
		IsReview       int64  `json:"is_review"`       // 是否审核
		CreatedAt      int64  `json:"created_at"`      // 评论时间
		UpdatedAt      int64  `json:"updated_at"`      // 更新时间
	}
)

func NewCommentModel(db *gorm.DB, cache *redis.Client) ICommentModel {
	return &defaultCommentModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameComment,
	}
}

// 切换事务操作
func (s *defaultCommentModel) WithTransaction(tx *gorm.DB) (out ICommentModel) {
	return NewCommentModel(tx, s.CacheEngin)
}

// 创建Comment记录
func (s *defaultCommentModel) Create(ctx context.Context, in *Comment) (out *Comment, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Comment记录
func (s *defaultCommentModel) Update(ctx context.Context, in *Comment) (out *Comment, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Comment记录
func (s *defaultCommentModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Comment{})
	return query.RowsAffected, query.Error
}

// 查询Comment记录
func (s *defaultCommentModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Comment)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Comment记录
func (s *defaultCommentModel) BatchCreate(ctx context.Context, in ...*Comment) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Comment记录
func (s *defaultCommentModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Comment{})
	return query.RowsAffected, query.Error
}

// 查询Comment总数
func (s *defaultCommentModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Comment{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Comment列表
func (s *defaultCommentModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Comment, err error) {
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

// 分页查询Comment记录
func (s *defaultCommentModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Comment, err error) {
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
