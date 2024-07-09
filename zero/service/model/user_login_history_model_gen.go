package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserLoginHistoryModel = (*defaultUserLoginHistoryModel)(nil)

type (
	// 接口定义
	UserLoginHistoryModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserLoginHistoryModel)
		// 插入
		Insert(ctx context.Context, in *UserLoginHistory) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*UserLoginHistory) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *UserLoginHistory) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *UserLoginHistory) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *UserLoginHistory, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserLoginHistory, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserLoginHistory, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserLoginHistory, err error)
		// add extra method in here
	}

	// 表字段定义
	UserLoginHistory struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // id
		UserId    int64     `json:"user_id" gorm:"column:user_id" `       // 用户id
		LoginType string    `json:"login_type" gorm:"column:login_type" ` // 登录类型
		Agent     string    `json:"agent" gorm:"column:agent" `           // 代理
		IpAddress string    `json:"ip_address" gorm:"column:ip_address" ` // ip host
		IpSource  string    `json:"ip_source" gorm:"column:ip_source" `   // ip 源
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultUserLoginHistoryModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewUserLoginHistoryModel(db *gorm.DB, cache *redis.Client) UserLoginHistoryModel {
	return &defaultUserLoginHistoryModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`user_login_history`",
	}
}

// 切换事务操作
func (m *defaultUserLoginHistoryModel) WithTransaction(tx *gorm.DB) (out UserLoginHistoryModel) {
	return NewUserLoginHistoryModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultUserLoginHistoryModel) Insert(ctx context.Context, in *UserLoginHistory) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultUserLoginHistoryModel) InsertBatch(ctx context.Context, in ...*UserLoginHistory) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultUserLoginHistoryModel) Update(ctx context.Context, in *UserLoginHistory) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultUserLoginHistoryModel) UpdateNotEmpty(ctx context.Context, in *UserLoginHistory) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultUserLoginHistoryModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&UserLoginHistory{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserLoginHistoryModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&UserLoginHistory{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserLoginHistoryModel) FindOne(ctx context.Context, id int64) (out *UserLoginHistory, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultUserLoginHistoryModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserLoginHistory, err error) {
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
func (m *defaultUserLoginHistoryModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询列表
func (m *defaultUserLoginHistoryModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserLoginHistory, err error) {
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
func (m *defaultUserLoginHistoryModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserLoginHistory, err error) {
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
