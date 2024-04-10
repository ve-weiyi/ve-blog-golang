package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUserAccount = "user_account"

type (
	// 接口定义
	UserAccountModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserAccountModel)
		// 增删改查
		Create(ctx context.Context, in *UserAccount) (out *UserAccount, err error)
		Update(ctx context.Context, in *UserAccount) (out *UserAccount, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserAccount, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UserAccount) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserAccount, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserAccount, err error)

		// 加载用户model
		LoadUserByUsername(ctx context.Context, username string) (data *UserAccount, err error)
	}

	// 接口实现
	defaultUserAccountModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UserAccount struct {
		Id           int64     `json:"id"`            // id
		Username     string    `json:"username"`      // 用户名
		Password     string    `json:"password"`      // 密码
		Status       int64     `json:"status"`        // 状态: -1删除 0正常 1禁用
		RegisterType string    `json:"register_type"` // 注册方式
		IpAddress    string    `json:"ip_address"`    // 注册ip
		IpSource     string    `json:"ip_source"`     // 注册ip 源
		CreatedAt    time.Time `json:"created_at"`    // 创建时间
		UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
	}
)

func NewUserAccountModel(db *gorm.DB, cache *redis.Client) UserAccountModel {
	return &defaultUserAccountModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUserAccount,
	}
}

// 切换事务操作
func (s *defaultUserAccountModel) WithTransaction(tx *gorm.DB) (out UserAccountModel) {
	return NewUserAccountModel(tx, s.CacheEngin)
}

// 创建UserAccount记录
func (s *defaultUserAccountModel) Create(ctx context.Context, in *UserAccount) (out *UserAccount, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UserAccount记录
func (s *defaultUserAccountModel) Update(ctx context.Context, in *UserAccount) (out *UserAccount, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UserAccount记录
func (s *defaultUserAccountModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UserAccount{})
	return query.RowsAffected, query.Error
}

// 查询UserAccount记录
func (s *defaultUserAccountModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserAccount, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UserAccount)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UserAccount记录
func (s *defaultUserAccountModel) BatchCreate(ctx context.Context, in ...*UserAccount) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UserAccount记录
func (s *defaultUserAccountModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UserAccount{})
	return query.RowsAffected, query.Error
}

// 查询UserAccount总数
func (s *defaultUserAccountModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UserAccount{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserAccount列表
func (s *defaultUserAccountModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserAccount, err error) {
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

// 分页查询UserAccount记录
func (s *defaultUserAccountModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserAccount, err error) {
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

// 加载用户model
func (s *defaultUserAccountModel) LoadUserByUsername(ctx context.Context, username string) (data *UserAccount, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	//查询用户信息
	err = db.Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
