package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleMenuModel = (*defaultRoleMenuModel)(nil)

type (
	// 接口定义
	RoleMenuModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleMenuModel)
		// 插入
		Insert(ctx context.Context, in *RoleMenu) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*RoleMenu) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *RoleMenu) (rows int64, err error)
		Save(ctx context.Context, in *RoleMenu) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *RoleMenu, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*RoleMenu, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*RoleMenu, err error)
		// add extra method in here
	}

	// 表字段定义
	RoleMenu struct {
		Id     int64 `json:"id" gorm:"column:id" `           // 主键id
		RoleId int64 `json:"role_id" gorm:"column:role_id" ` // 角色id
		MenuId int64 `json:"menu_id" gorm:"column:menu_id" ` // 菜单id
	}

	// 接口实现
	defaultRoleMenuModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewRoleMenuModel(db *gorm.DB, cache *redis.Client) RoleMenuModel {
	return &defaultRoleMenuModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`role_menu`",
	}
}

// 切换事务操作
func (m *defaultRoleMenuModel) WithTransaction(tx *gorm.DB) (out RoleMenuModel) {
	return NewRoleMenuModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultRoleMenuModel) Insert(ctx context.Context, in *RoleMenu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultRoleMenuModel) InsertBatch(ctx context.Context, in ...*RoleMenu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultRoleMenuModel) Update(ctx context.Context, in *RoleMenu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultRoleMenuModel) Save(ctx context.Context, in *RoleMenu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultRoleMenuModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&RoleMenu{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleMenuModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&RoleMenu{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleMenuModel) FindOne(ctx context.Context, id int64) (out *RoleMenu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultRoleMenuModel) First(ctx context.Context, conditions string, args ...interface{}) (out *RoleMenu, err error) {
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
func (m *defaultRoleMenuModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&RoleMenu{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultRoleMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleMenu, err error) {
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
func (m *defaultRoleMenuModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*RoleMenu, err error) {
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
