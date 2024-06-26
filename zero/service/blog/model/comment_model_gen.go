// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cacheCommentIdPrefix = "cache:comment:id:"
)

type (
	commentModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out CommentModel)
		Insert(ctx context.Context, in *Comment) (*Comment, error)
		InsertBatch(ctx context.Context, in ...*Comment) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Comment, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (out []*Comment, err error)
		FindOne(ctx context.Context, id int64) (*Comment, error)
		Update(ctx context.Context, data *Comment) (*Comment, error)
		UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error)
		Save(ctx context.Context, data *Comment) (*Comment, error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// add extra method in here

	}

	defaultCommentModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}

	Comment struct {
		Id             int64     `gorm:"column:id"`              // 主键
		TopicId        int64     `gorm:"column:topic_id"`        // 主题id
		ParentId       int64     `gorm:"column:parent_id"`       // 父评论id
		SessionId      int64     `gorm:"column:session_id"`      // 会话id
		UserId         int64     `gorm:"column:user_id"`         // 评论用户id
		ReplyUserId    int64     `gorm:"column:reply_user_id"`   // 评论回复用户id
		CommentContent string    `gorm:"column:comment_content"` // 评论内容
		LikeCount      int64     `gorm:"column:like_count"`      // 评论点赞数量
		Type           int64     `gorm:"column:type"`            // 评论类型 1.文章 2.友链 3.说说
		Status         int64     `gorm:"column:status"`          // 状态 0.正常 1.已编辑 2.已删除
		IsReview       int64     `gorm:"column:is_review"`       // 是否审核
		CreatedAt      time.Time `gorm:"column:created_at"`      // 创建时间
		UpdatedAt      time.Time `gorm:"column:updated_at"`      // 更新时间
	}
)

func newCommentModel(db *gorm.DB, cache *redis.Client) *defaultCommentModel {
	return &defaultCommentModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`comment`",
	}
}

// 切换事务操作
func (m *defaultCommentModel) WithTransaction(tx *gorm.DB) (out CommentModel) {
	return NewCommentModel(tx, m.CacheEngin)
}

// 插入Comment记录
func (m *defaultCommentModel) Insert(ctx context.Context, in *Comment) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 插入Comment记录
func (m *defaultCommentModel) InsertBatch(ctx context.Context, in ...*Comment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 更新Comment记录
func (m *defaultCommentModel) Update(ctx context.Context, in *Comment) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Updates(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Comment记录
func (m *defaultCommentModel) UpdateColumns(ctx context.Context, id int64, columns map[string]interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where("`id` = ?", id).UpdateColumns(&columns)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

// 更新Comment记录
func (m *defaultCommentModel) Save(ctx context.Context, in *Comment) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Comment记录
func (m *defaultCommentModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	query := db.Delete(&Comment{})
	return query.RowsAffected, query.Error
}

// 删除Comment记录
func (m *defaultCommentModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Comment{})
	return result.RowsAffected, result.Error
}

// 查询Comment记录
func (m *defaultCommentModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询Comment总数
func (m *defaultCommentModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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
func (m *defaultCommentModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Comment, err error) {
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

// 分页查询Comment记录
func (m *defaultCommentModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Comment, err error) {
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
	if limit > 0 || offset > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询Comment记录
func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

func (m *defaultCommentModel) TableName() string {
	return m.table
}
