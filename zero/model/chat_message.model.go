package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameChatMessage = "chat_message"

type (
	// 接口定义
	ChatMessageModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ChatMessageModel)
		// 增删改查
		Create(ctx context.Context, in *ChatMessage) (out *ChatMessage, err error)
		Update(ctx context.Context, in *ChatMessage) (out *ChatMessage, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ChatMessage, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*ChatMessage) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*ChatMessage, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatMessage, err error)
	}

	// 接口实现
	defaultChatMessageModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	ChatMessage struct {
		Id         int64     `json:"id"`           // 主键
		ChatId     string    `json:"chat_id"`      // 聊天id
		UserId     int64     `json:"user_id"`      // 用户id
		ReplyMsgId int64     `json:"reply_msg_id"` // 回复消息id
		Content    string    `json:"content"`      // 聊天内容
		IpAddress  string    `json:"ip_address"`   // ip地址
		IpSource   string    `json:"ip_source"`    // ip来源
		Type       int64     `json:"type"`         // 类型
		Status     int64     `json:"status"`       // 0正常 1撤回 2已编辑
		CreatedAt  time.Time `json:"created_at"`   // 创建时间
		UpdatedAt  time.Time `json:"updated_at"`   // 更新时间
	}
)

func NewChatMessageModel(db *gorm.DB, cache *redis.Client) ChatMessageModel {
	return &defaultChatMessageModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameChatMessage,
	}
}

// 切换事务操作
func (s *defaultChatMessageModel) WithTransaction(tx *gorm.DB) (out ChatMessageModel) {
	return NewChatMessageModel(tx, s.CacheEngin)
}

// 创建ChatMessage记录
func (s *defaultChatMessageModel) Create(ctx context.Context, in *ChatMessage) (out *ChatMessage, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新ChatMessage记录
func (s *defaultChatMessageModel) Update(ctx context.Context, in *ChatMessage) (out *ChatMessage, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除ChatMessage记录
func (s *defaultChatMessageModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&ChatMessage{})
	return query.RowsAffected, query.Error
}

// 查询ChatMessage记录
func (s *defaultChatMessageModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ChatMessage, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(ChatMessage)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询ChatMessage记录
func (s *defaultChatMessageModel) BatchCreate(ctx context.Context, in ...*ChatMessage) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询ChatMessage记录
func (s *defaultChatMessageModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&ChatMessage{})
	return query.RowsAffected, query.Error
}

// 查询ChatMessage总数
func (s *defaultChatMessageModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询ChatMessage列表
func (s *defaultChatMessageModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ChatMessage, err error) {
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

// 分页查询ChatMessage记录
func (s *defaultChatMessageModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatMessage, err error) {
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
