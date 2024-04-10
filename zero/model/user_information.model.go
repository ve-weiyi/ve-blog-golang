package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUserInformation = "user_information"

type (
	// 接口定义
	UserInformationModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserInformationModel)
		// 增删改查
		Create(ctx context.Context, in *UserInformation) (out *UserInformation, err error)
		Update(ctx context.Context, in *UserInformation) (out *UserInformation, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserInformation, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UserInformation) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserInformation, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserInformation, err error)

		FindUserInfo(ctx context.Context, userId int64) (out *UserInformation, err error)
	}

	// 接口实现
	defaultUserInformationModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UserInformation struct {
		Id        int64     `json:"id"`         // id
		UserId    int64     `json:"user_id"`    // 用户id
		Email     string    `json:"email"`      // 用户邮箱
		Nickname  string    `json:"nickname"`   // 用户昵称
		Avatar    string    `json:"avatar"`     // 用户头像
		Phone     string    `json:"phone"`      // 用户手机号
		Intro     string    `json:"intro"`      // 个人简介
		Website   string    `json:"website"`    // 个人网站
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewUserInformationModel(db *gorm.DB, cache *redis.Client) UserInformationModel {
	return &defaultUserInformationModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUserInformation,
	}
}

// 切换事务操作
func (s *defaultUserInformationModel) WithTransaction(tx *gorm.DB) (out UserInformationModel) {
	return NewUserInformationModel(tx, s.CacheEngin)
}

// 创建UserInformation记录
func (s *defaultUserInformationModel) Create(ctx context.Context, in *UserInformation) (out *UserInformation, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UserInformation记录
func (s *defaultUserInformationModel) Update(ctx context.Context, in *UserInformation) (out *UserInformation, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UserInformation记录
func (s *defaultUserInformationModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UserInformation{})
	return query.RowsAffected, query.Error
}

// 查询UserInformation记录
func (s *defaultUserInformationModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserInformation, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UserInformation)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UserInformation记录
func (s *defaultUserInformationModel) BatchCreate(ctx context.Context, in ...*UserInformation) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UserInformation记录
func (s *defaultUserInformationModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UserInformation{})
	return query.RowsAffected, query.Error
}

// 查询UserInformation总数
func (s *defaultUserInformationModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UserInformation{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserInformation列表
func (s *defaultUserInformationModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserInformation, err error) {
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

// 分页查询UserInformation记录
func (s *defaultUserInformationModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UserInformation, err error) {
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

// 查询用户信息
func (s *defaultUserInformationModel) FindUserInfo(ctx context.Context, userId int64) (out *UserInformation, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	//查询用户信息
	err = db.Where("user_id = ?", userId).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
