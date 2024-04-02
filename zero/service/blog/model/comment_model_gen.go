package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ CommentModel = (*defaultCommentModel)(nil)

type (
	// 接口定义
	CommentModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out CommentModel)
		// 插入
		Insert(ctx context.Context, in *Comment) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Comment) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *Comment) (rows int64, err error)
		Save(ctx context.Context, in *Comment) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Comment, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Comment, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Comment, err error)
		// add extra method in here
	}

	// 表字段定义
	Comment struct {
		Id             int64     `json:"id" gorm:"column:id" `                           // 主键
		TopicId        int64     `json:"topic_id" gorm:"column:topic_id" `               // 主题id
		ParentId       int64     `json:"parent_id" gorm:"column:parent_id" `             // 父评论id
		SessionId      int64     `json:"session_id" gorm:"column:session_id" `           // 会话id
		UserId         int64     `json:"user_id" gorm:"column:user_id" `                 // 评论用户id
		ReplyUserId    int64     `json:"reply_user_id" gorm:"column:reply_user_id" `     // 评论回复用户id
		CommentContent string    `json:"comment_content" gorm:"column:comment_content" ` // 评论内容
		LikeCount      int64     `json:"like_count" gorm:"column:like_count" `           // 评论点赞数量
		Type           int64     `json:"type" gorm:"column:type" `                       // 评论类型 1.文章 2.友链 3.说说
		Status         int64     `json:"status" gorm:"column:status" `                   // 状态 0.正常 1.已编辑 2.已删除
		IsReview       int64     `json:"is_review" gorm:"column:is_review" `             // 是否审核
		CreatedAt      time.Time `json:"created_at" gorm:"column:created_at" `           // 创建时间
		UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at" `           // 更新时间
	}

	// 接口实现
	defaultCommentModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewCommentModel(db *gorm.DB, cache *redis.Client) CommentModel {
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

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultCommentModel) Insert(ctx context.Context, in *Comment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultCommentModel) InsertBatch(ctx context.Context, in ...*Comment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultCommentModel) Update(ctx context.Context, in *Comment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultCommentModel) Save(ctx context.Context, in *Comment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultCommentModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Comment{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultCommentModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Comment{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (out *Comment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultCommentModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Comment, err error) {
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

// 查询列表
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

// 分页查询记录
func (m *defaultCommentModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Comment, err error) {
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
