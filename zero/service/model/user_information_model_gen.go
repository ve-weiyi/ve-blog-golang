package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserInformationModel = (*defaultUserInformationModel)(nil)

type (
	// 接口定义
	UserInformationModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserInformationModel)
		// 插入
		Insert(ctx context.Context, in *UserInformation) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*UserInformation) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *UserInformation) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *UserInformation) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *UserInformation, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserInformation, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserInformation, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserInformation, err error)
		// add extra method in here
		FindOneByUserId(ctx context.Context, user_id int64) (out *UserInformation, err error)
	}

	// 表字段定义
	UserInformation struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // id
		UserId    int64     `json:"user_id" gorm:"column:user_id" `       // 用户id
		Email     string    `json:"email" gorm:"column:email" `           // 用户邮箱
		Nickname  string    `json:"nickname" gorm:"column:nickname" `     // 用户昵称
		Avatar    string    `json:"avatar" gorm:"column:avatar" `         // 用户头像
		Phone     string    `json:"phone" gorm:"column:phone" `           // 用户手机号
		Intro     string    `json:"intro" gorm:"column:intro" `           // 个人简介
		Website   string    `json:"website" gorm:"column:website" `       // 个人网站
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultUserInformationModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewUserInformationModel(db *gorm.DB, cache *redis.Client) UserInformationModel {
	return &defaultUserInformationModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`user_information`",
	}
}

// 切换事务操作
func (m *defaultUserInformationModel) WithTransaction(tx *gorm.DB) (out UserInformationModel) {
	return NewUserInformationModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultUserInformationModel) Insert(ctx context.Context, in *UserInformation) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultUserInformationModel) InsertBatch(ctx context.Context, in ...*UserInformation) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultUserInformationModel) Update(ctx context.Context, in *UserInformation) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultUserInformationModel) UpdateNotEmpty(ctx context.Context, in *UserInformation) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultUserInformationModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&UserInformation{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserInformationModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&UserInformation{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserInformationModel) FindOne(ctx context.Context, id int64) (out *UserInformation, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultUserInformationModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserInformation, err error) {
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
func (m *defaultUserInformationModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询列表
func (m *defaultUserInformationModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserInformation, err error) {
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
func (m *defaultUserInformationModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserInformation, err error) {
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
func (m *defaultUserInformationModel) FindOneByUserId(ctx context.Context, user_id int64) (out *UserInformation, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`user_id` = ?", user_id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
