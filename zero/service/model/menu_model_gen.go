package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ MenuModel = (*defaultMenuModel)(nil)

type (
	// 接口定义
	MenuModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out MenuModel)
		// 插入
		Insert(ctx context.Context, in *Menu) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Menu) (rows int64, err error)
		// 更新
		Save(ctx context.Context, in *Menu) (rows int64, err error)
		Update(ctx context.Context, in *Menu) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Menu, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Menu, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Menu, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Menu, err error)
		// add extra method in here
		FindOneByPath(ctx context.Context, path string) (out *Menu, err error)
	}

	// 表字段定义
	Menu struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // 主键
		ParentId  int64     `json:"parent_id" gorm:"column:parent_id" `   // 父id
		Title     string    `json:"title" gorm:"column:title" `           // 菜单标题
		Path      string    `json:"path" gorm:"column:path" `             // 路由路径
		Name      string    `json:"name" gorm:"column:name" `             // 路由名称
		Component string    `json:"component" gorm:"column:component" `   // 路由组件
		Redirect  string    `json:"redirect" gorm:"column:redirect" `     // 路由重定向
		Type      int64     `json:"type" gorm:"column:type" `             // 菜单类型
		Rank      int64     `json:"rank" gorm:"column:rank" `             // 排序
		Extra     string    `json:"extra" gorm:"column:extra" `           // 菜单元数据
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultMenuModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewMenuModel(db *gorm.DB, cache *redis.Client) MenuModel {
	return &defaultMenuModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`menu`",
	}
}

// 切换事务操作
func (m *defaultMenuModel) WithTransaction(tx *gorm.DB) (out MenuModel) {
	return NewMenuModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultMenuModel) Insert(ctx context.Context, in *Menu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultMenuModel) InsertBatch(ctx context.Context, in ...*Menu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultMenuModel) Save(ctx context.Context, in *Menu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultMenuModel) Update(ctx context.Context, in *Menu) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultMenuModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Menu{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultMenuModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Menu{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultMenuModel) FindOne(ctx context.Context, id int64) (out *Menu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultMenuModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Menu, err error) {
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
func (m *defaultMenuModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Menu{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Menu, err error) {
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
func (m *defaultMenuModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Menu, err error) {
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
func (m *defaultMenuModel) FindOneByPath(ctx context.Context, path string) (out *Menu, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`path` = ?", path).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
