package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameChatSession = "chat_session"

type (
	// 接口定义
	ChatSessionModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ChatSessionModel)
		// 增删改查
		Create(ctx context.Context, in *ChatSession) (out *ChatSession, err error)
		Update(ctx context.Context, in *ChatSession) (out *ChatSession, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ChatSession, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*ChatSession) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*ChatSession, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatSession, err error)
	}

	// 接口实现
	defaultChatSessionModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	ChatSession struct {
		Id        int64     `json:"id"`         // 主键
		ChatId    string    `json:"chat_id"`    // 聊天id
		ChatTitle string    `json:"chat_title"` // 标题
		Type      string    `json:"type"`       // 类型
		Status    int64     `json:"status"`     // 0正常 1删除
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewChatSessionModel(db *gorm.DB, cache *redis.Client) ChatSessionModel {
	return &defaultChatSessionModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameChatSession,
	}
}

// 切换事务操作
func (s *defaultChatSessionModel) WithTransaction(tx *gorm.DB) (out ChatSessionModel) {
	return NewChatSessionModel(tx, s.CacheEngin)
}

// 创建ChatSession记录
func (s *defaultChatSessionModel) Create(ctx context.Context, in *ChatSession) (out *ChatSession, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新ChatSession记录
func (s *defaultChatSessionModel) Update(ctx context.Context, in *ChatSession) (out *ChatSession, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除ChatSession记录
func (s *defaultChatSessionModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&ChatSession{})
	return query.RowsAffected, query.Error
}

// 查询ChatSession记录
func (s *defaultChatSessionModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ChatSession, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(ChatSession)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询ChatSession记录
func (s *defaultChatSessionModel) BatchCreate(ctx context.Context, in ...*ChatSession) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询ChatSession记录
func (s *defaultChatSessionModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&ChatSession{})
	return query.RowsAffected, query.Error
}

// 查询ChatSession总数
func (s *defaultChatSessionModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&ChatSession{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询ChatSession列表
func (s *defaultChatSessionModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ChatSession, err error) {
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

// 分页查询ChatSession记录
func (s *defaultChatSessionModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatSession, err error) {
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
