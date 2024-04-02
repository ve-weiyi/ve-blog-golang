package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleModel = (*defaultRoleModel)(nil)

type (
	// 接口定义
	RoleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleModel)
		// 插入
		Insert(ctx context.Context, in *Role) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Role) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *Role) (rows int64, err error)
		Save(ctx context.Context, in *Role) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Role, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Role, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Role, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Role, err error)
		// add extra method in here
	}

	// 表字段定义
	Role struct {
		Id          int64     `json:"id" gorm:"column:id" `                     // 主键id
		ParentId    int64     `json:"parent_id" gorm:"column:parent_id" `       // 父角色id
		RoleDomain  string    `json:"role_domain" gorm:"column:role_domain" `   // 角色域
		RoleName    string    `json:"role_name" gorm:"column:role_name" `       // 角色名
		RoleComment string    `json:"role_comment" gorm:"column:role_comment" ` // 角色备注
		IsDisable   int64     `json:"is_disable" gorm:"column:is_disable" `     // 是否禁用  0否 1是
		IsDefault   int64     `json:"is_default" gorm:"column:is_default" `     // 是否默认角色 0否 1是
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at" `     // 创建时间
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at" `     // 更新时间
	}

	// 接口实现
	defaultRoleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewRoleModel(db *gorm.DB, cache *redis.Client) RoleModel {
	return &defaultRoleModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`role`",
	}
}

// 切换事务操作
func (m *defaultRoleModel) WithTransaction(tx *gorm.DB) (out RoleModel) {
	return NewRoleModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultRoleModel) Insert(ctx context.Context, in *Role) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultRoleModel) InsertBatch(ctx context.Context, in ...*Role) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultRoleModel) Update(ctx context.Context, in *Role) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultRoleModel) Save(ctx context.Context, in *Role) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultRoleModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Role{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Role{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleModel) FindOne(ctx context.Context, id int64) (out *Role, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultRoleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Role, err error) {
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
func (m *defaultRoleModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Role{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultRoleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Role, err error) {
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
func (m *defaultRoleModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Role, err error) {
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
