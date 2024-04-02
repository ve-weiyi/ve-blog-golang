package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameMenu = "menu"

type (
	// 接口定义
	IMenuModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out IMenuModel)
		// 增删改查
		Create(ctx context.Context, in *Menu) (out *Menu, err error)
		Update(ctx context.Context, in *Menu) (out *Menu, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Menu, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Menu) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Menu, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Menu, err error)
	}

	// 接口实现
	defaultMenuModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Menu struct {
		ID        int64  `json:"id"`         // 主键
		ParentID  int64  `json:"parent_id"`  // 父id
		Title     string `json:"title"`      // 菜单标题
		Path      string `json:"path"`       // 路由路径
		Name      string `json:"name"`       // 路由名称
		Component string `json:"component"`  // 路由组件
		Redirect  string `json:"redirect"`   // 路由重定向
		Type      int64  `json:"type"`       // 菜单类型
		Meta      string `json:"meta"`       // 菜单元数据
		CreatedAt int64  `json:"created_at"` // 创建时间
		UpdatedAt int64  `json:"updated_at"` // 更新时间
	}
)

func NewMenuModel(db *gorm.DB, cache *redis.Client) IMenuModel {
	return &defaultMenuModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameMenu,
	}
}

// 切换事务操作
func (s *defaultMenuModel) WithTransaction(tx *gorm.DB) (out IMenuModel) {
	return NewMenuModel(tx, s.CacheEngin)
}

// 创建Menu记录
func (s *defaultMenuModel) Create(ctx context.Context, in *Menu) (out *Menu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Menu记录
func (s *defaultMenuModel) Update(ctx context.Context, in *Menu) (out *Menu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Menu记录
func (s *defaultMenuModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Menu{})
	return query.RowsAffected, query.Error
}

// 查询Menu记录
func (s *defaultMenuModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Menu, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Menu)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Menu记录
func (s *defaultMenuModel) BatchCreate(ctx context.Context, in ...*Menu) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Menu记录
func (s *defaultMenuModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Menu{})
	return query.RowsAffected, query.Error
}

// 查询Menu总数
func (s *defaultMenuModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询Menu列表
func (s *defaultMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Menu, err error) {
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

// 分页查询Menu记录
func (s *defaultMenuModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Menu, err error) {
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
