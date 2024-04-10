package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUserOauth = "user_oauth"

type (
	// 接口定义
	UserOauthModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserOauthModel)
		// 增删改查
		Create(ctx context.Context, in *UserOauth) (out *UserOauth, err error)
		Update(ctx context.Context, in *UserOauth) (out *UserOauth, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserOauth, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UserOauth) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserOauth, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserOauth, err error)
	}

	// 接口实现
	defaultUserOauthModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UserOauth struct {
		Id        int64     `json:"id"`         // id
		UserId    int64     `json:"user_id"`    // 用户id
		OpenId    string    `json:"open_id"`    // 开发平台id，标识唯一用户
		Platform  string    `json:"platform"`   // 平台:手机号、邮箱、微信、飞书
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewUserOauthModel(db *gorm.DB, cache *redis.Client) UserOauthModel {
	return &defaultUserOauthModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUserOauth,
	}
}

// 切换事务操作
func (s *defaultUserOauthModel) WithTransaction(tx *gorm.DB) (out UserOauthModel) {
	return NewUserOauthModel(tx, s.CacheEngin)
}

// 创建UserOauth记录
func (s *defaultUserOauthModel) Create(ctx context.Context, in *UserOauth) (out *UserOauth, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UserOauth记录
func (s *defaultUserOauthModel) Update(ctx context.Context, in *UserOauth) (out *UserOauth, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UserOauth记录
func (s *defaultUserOauthModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UserOauth{})
	return query.RowsAffected, query.Error
}

// 查询UserOauth记录
func (s *defaultUserOauthModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserOauth, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UserOauth)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UserOauth记录
func (s *defaultUserOauthModel) BatchCreate(ctx context.Context, in ...*UserOauth) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UserOauth记录
func (s *defaultUserOauthModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UserOauth{})
	return query.RowsAffected, query.Error
}

// 查询UserOauth总数
func (s *defaultUserOauthModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UserOauth{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserOauth列表
func (s *defaultUserOauthModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserOauth, err error) {
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

// 分页查询UserOauth记录
func (s *defaultUserOauthModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserOauth, err error) {
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
