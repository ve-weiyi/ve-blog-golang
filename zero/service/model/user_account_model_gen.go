package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ UserAccountModel = (*defaultUserAccountModel)(nil)

type (
	// 接口定义
	UserAccountModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UserAccountModel)
		// 插入
		Insert(ctx context.Context, in *UserAccount) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*UserAccount) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *UserAccount) (rows int64, err error)
		Save(ctx context.Context, in *UserAccount) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *UserAccount, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UserAccount, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UserAccount, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserAccount, err error)
		// add extra method in here
		FindOneByUsername(ctx context.Context, username string) (out *UserAccount, err error)
	}

	// 表字段定义
	UserAccount struct {
		Id           int64     `json:"id" gorm:"column:id" `                       // id
		Username     string    `json:"username" gorm:"column:username" `           // 用户名
		Password     string    `json:"password" gorm:"column:password" `           // 密码
		Status       int64     `json:"status" gorm:"column:status" `               // 状态: -1删除 0正常 1禁用
		RegisterType string    `json:"register_type" gorm:"column:register_type" ` // 注册方式
		IpAddress    string    `json:"ip_address" gorm:"column:ip_address" `       // 注册ip
		IpSource     string    `json:"ip_source" gorm:"column:ip_source" `         // 注册ip 源
		LoginAt      time.Time `json:"login_at" gorm:"column:login_at" `           // 登录时间
		LogoutAt     time.Time `json:"logout_at" gorm:"column:logout_at" `         // 登出时间
		CreatedAt    time.Time `json:"created_at" gorm:"column:created_at" `       // 创建时间
		UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at" `       // 更新时间
	}

	// 接口实现
	defaultUserAccountModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewUserAccountModel(db *gorm.DB, cache *redis.Client) UserAccountModel {
	return &defaultUserAccountModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`user_account`",
	}
}

// 切换事务操作
func (m *defaultUserAccountModel) WithTransaction(tx *gorm.DB) (out UserAccountModel) {
	return NewUserAccountModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultUserAccountModel) Insert(ctx context.Context, in *UserAccount) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultUserAccountModel) InsertBatch(ctx context.Context, in ...*UserAccount) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultUserAccountModel) Update(ctx context.Context, in *UserAccount) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultUserAccountModel) Save(ctx context.Context, in *UserAccount) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultUserAccountModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&UserAccount{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserAccountModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&UserAccount{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultUserAccountModel) FindOne(ctx context.Context, id int64) (out *UserAccount, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultUserAccountModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UserAccount, err error) {
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
func (m *defaultUserAccountModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询列表
func (m *defaultUserAccountModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UserAccount, err error) {
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
func (m *defaultUserAccountModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*UserAccount, err error) {
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
func (m *defaultUserAccountModel) FindOneByUsername(ctx context.Context, username string) (out *UserAccount, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`username` = ?", username).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
