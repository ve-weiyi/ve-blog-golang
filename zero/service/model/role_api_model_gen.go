package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ RoleApiModel = (*defaultRoleApiModel)(nil)

type (
	// 接口定义
	RoleApiModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RoleApiModel)
		// 插入
		Insert(ctx context.Context, in *RoleApi) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*RoleApi) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *RoleApi) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *RoleApi) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *RoleApi, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *RoleApi, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*RoleApi, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*RoleApi, err error)
		// add extra method in here
	}

	// 表字段定义
	RoleApi struct {
		Id     int64 `json:"id" gorm:"column:id" `           // 主键id
		RoleId int64 `json:"role_id" gorm:"column:role_id" ` // 角色id
		ApiId  int64 `json:"api_id" gorm:"column:api_id" `   // 接口id
	}

	// 接口实现
	defaultRoleApiModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewRoleApiModel(db *gorm.DB, cache *redis.Client) RoleApiModel {
	return &defaultRoleApiModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`role_api`",
	}
}

// 切换事务操作
func (m *defaultRoleApiModel) WithTransaction(tx *gorm.DB) (out RoleApiModel) {
	return NewRoleApiModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultRoleApiModel) Insert(ctx context.Context, in *RoleApi) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultRoleApiModel) InsertBatch(ctx context.Context, in ...*RoleApi) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultRoleApiModel) Update(ctx context.Context, in *RoleApi) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultRoleApiModel) UpdateNotEmpty(ctx context.Context, in *RoleApi) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultRoleApiModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&RoleApi{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleApiModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&RoleApi{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultRoleApiModel) FindOne(ctx context.Context, id int64) (out *RoleApi, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultRoleApiModel) First(ctx context.Context, conditions string, args ...interface{}) (out *RoleApi, err error) {
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
func (m *defaultRoleApiModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&RoleApi{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultRoleApiModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*RoleApi, err error) {
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
func (m *defaultRoleApiModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*RoleApi, err error) {
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
