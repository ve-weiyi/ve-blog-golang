package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ TCommentModel = (*defaultTCommentModel)(nil)

type (
	// 接口定义
	TCommentModel interface {
		TableName() string
		// 在事务中操作
		WithTransaction(tx *gorm.DB) (out TCommentModel)
		// 插入
		Insert(ctx context.Context, in *TComment) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*TComment) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *TComment) (rows int64, err error)
		Save(ctx context.Context, in *TComment) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *TComment, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *TComment, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*TComment, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TComment, err error)
		// add extra method in here
	}

	// 表字段定义
	TComment struct {
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
	defaultTCommentModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewTCommentModel(db *gorm.DB, cache *redis.Client) TCommentModel {
	return &defaultTCommentModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`t_comment`",
	}
}

func (m *defaultTCommentModel) TableName() string {
	return m.table
}

// 在事务中操作
func (m *defaultTCommentModel) WithTransaction(tx *gorm.DB) (out TCommentModel) {
	return NewTCommentModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultTCommentModel) Insert(ctx context.Context, in *TComment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultTCommentModel) InsertBatch(ctx context.Context, in ...*TComment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultTCommentModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&TComment{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultTCommentModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&TComment{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 保存记录（更新零值）
func (m *defaultTCommentModel) Save(ctx context.Context, in *TComment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultTCommentModel) Update(ctx context.Context, in *TComment) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultTCommentModel) FindOne(ctx context.Context, id int64) (out *TComment, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultTCommentModel) First(ctx context.Context, conditions string, args ...interface{}) (out *TComment, err error) {
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
func (m *defaultTCommentModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&TComment{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultTCommentModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*TComment, err error) {
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
func (m *defaultTCommentModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TComment, err error) {
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
