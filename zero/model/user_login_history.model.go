package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUserLoginHistory = "user_login_history"

type (
	// 接口定义
	UserLoginHistoryModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserLoginHistoryModel)
		// 增删改查
		Create(ctx context.Context, in *UserLoginHistory) (out *UserLoginHistory, err error)
		Update(ctx context.Context, in *UserLoginHistory) (out *UserLoginHistory, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserLoginHistory, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UserLoginHistory) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserLoginHistory, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserLoginHistory, err error)
	}

	// 接口实现
	defaultUserLoginHistoryModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UserLoginHistory struct {
		Id        int64     `json:"id"`         // id
		UserId    int64     `json:"user_id"`    // 用户id
		LoginType string    `json:"login_type"` // 登录类型
		Agent     string    `json:"agent"`      // 代理
		IpAddress string    `json:"ip_address"` // ip host
		IpSource  string    `json:"ip_source"`  // ip 源
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewUserLoginHistoryModel(db *gorm.DB, cache *redis.Client) UserLoginHistoryModel {
	return &defaultUserLoginHistoryModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUserLoginHistory,
	}
}

// 切换事务操作
func (s *defaultUserLoginHistoryModel) WithTransaction(tx *gorm.DB) (out UserLoginHistoryModel) {
	return NewUserLoginHistoryModel(tx, s.CacheEngin)
}

// 创建UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) Create(ctx context.Context, in *UserLoginHistory) (out *UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) Update(ctx context.Context, in *UserLoginHistory) (out *UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UserLoginHistory{})
	return query.RowsAffected, query.Error
}

// 查询UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UserLoginHistory)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) BatchCreate(ctx context.Context, in ...*UserLoginHistory) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UserLoginHistory{})
	return query.RowsAffected, query.Error
}

// 查询UserLoginHistory总数
func (s *defaultUserLoginHistoryModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UserLoginHistory{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserLoginHistory列表
func (s *defaultUserLoginHistoryModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserLoginHistory, err error) {
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

// 分页查询UserLoginHistory记录
func (s *defaultUserLoginHistoryModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserLoginHistory, err error) {
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
