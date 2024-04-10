package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameChatRecord = "chat_record"

type (
	// 接口定义
	ChatRecordModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ChatRecordModel)
		// 增删改查
		Create(ctx context.Context, in *ChatRecord) (out *ChatRecord, err error)
		Update(ctx context.Context, in *ChatRecord) (out *ChatRecord, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *ChatRecord, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*ChatRecord) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*ChatRecord, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatRecord, err error)
	}

	// 接口实现
	defaultChatRecordModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	ChatRecord struct {
		Id        int64     `json:"id"`         // 主键
		UserId    int64     `json:"user_id"`    // 用户id
		Nickname  string    `json:"nickname"`   // 昵称
		Avatar    string    `json:"avatar"`     // 头像
		Content   string    `json:"content"`    // 聊天内容
		IpAddress string    `json:"ip_address"` // ip地址
		IpSource  string    `json:"ip_source"`  // ip来源
		Type      int64     `json:"type"`       // 类型
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewChatRecordModel(db *gorm.DB, cache *redis.Client) ChatRecordModel {
	return &defaultChatRecordModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameChatRecord,
	}
}

// 切换事务操作
func (s *defaultChatRecordModel) WithTransaction(tx *gorm.DB) (out ChatRecordModel) {
	return NewChatRecordModel(tx, s.CacheEngin)
}

// 创建ChatRecord记录
func (s *defaultChatRecordModel) Create(ctx context.Context, in *ChatRecord) (out *ChatRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新ChatRecord记录
func (s *defaultChatRecordModel) Update(ctx context.Context, in *ChatRecord) (out *ChatRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除ChatRecord记录
func (s *defaultChatRecordModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&ChatRecord{})
	return query.RowsAffected, query.Error
}

// 查询ChatRecord记录
func (s *defaultChatRecordModel) First(ctx context.Context, conditions string, args ...interface{}) (out *ChatRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(ChatRecord)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询ChatRecord记录
func (s *defaultChatRecordModel) BatchCreate(ctx context.Context, in ...*ChatRecord) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询ChatRecord记录
func (s *defaultChatRecordModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&ChatRecord{})
	return query.RowsAffected, query.Error
}

// 查询ChatRecord总数
func (s *defaultChatRecordModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&ChatRecord{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询ChatRecord列表
func (s *defaultChatRecordModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*ChatRecord, err error) {
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

// 分页查询ChatRecord记录
func (s *defaultChatRecordModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*ChatRecord, err error) {
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
