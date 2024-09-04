package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ChatMessageModel = (*defaultChatMessageModel)(nil)

type (
	// 接口定义
	ChatMessageModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ChatMessageModel)
		// 插入
		Insert(ctx context.Context, in *ChatMessage) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*ChatMessage) (rows int64, err error)
		// 更新
		Save(ctx context.Context, in *ChatMessage) (rows int64, err error)
		Update(ctx context.Context, in *ChatMessage) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *ChatMessage, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ChatMessage, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*ChatMessage, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*ChatMessage, err error)
		// add extra method in here
	}

	// 表字段定义
	ChatMessage struct {
		Id         int64     `json:"id" gorm:"column:id" `                     // 主键
		ChatId     string    `json:"chat_id" gorm:"column:chat_id" `           // 聊天id
		UserId     int64     `json:"user_id" gorm:"column:user_id" `           // 用户id
		ReplyMsgId int64     `json:"reply_msg_id" gorm:"column:reply_msg_id" ` // 回复消息id
		Content    string    `json:"content" gorm:"column:content" `           // 聊天内容
		IpAddress  string    `json:"ip_address" gorm:"column:ip_address" `     // ip地址
		IpSource   string    `json:"ip_source" gorm:"column:ip_source" `       // ip来源
		Type       int64     `json:"type" gorm:"column:type" `                 // 类型
		Status     int64     `json:"status" gorm:"column:status" `             // 0正常 1撤回 2已编辑
		CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" `     // 创建时间
		UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" `     // 更新时间
	}

	// 接口实现
	defaultChatMessageModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewChatMessageModel(db *gorm.DB, cache *redis.Client) ChatMessageModel {
	return &defaultChatMessageModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`chat_message`",
	}
}

// 切换事务操作
func (m *defaultChatMessageModel) WithTransaction(tx *gorm.DB) (out ChatMessageModel) {
	return NewChatMessageModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultChatMessageModel) Insert(ctx context.Context, in *ChatMessage) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultChatMessageModel) InsertBatch(ctx context.Context, in ...*ChatMessage) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultChatMessageModel) Save(ctx context.Context, in *ChatMessage) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultChatMessageModel) Update(ctx context.Context, in *ChatMessage) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultChatMessageModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&ChatMessage{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultChatMessageModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&ChatMessage{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultChatMessageModel) FindOne(ctx context.Context, id int64) (out *ChatMessage, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultChatMessageModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ChatMessage, err error) {
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
func (m *defaultChatMessageModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&ChatMessage{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultChatMessageModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ChatMessage, err error) {
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
func (m *defaultChatMessageModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*ChatMessage, err error) {
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
