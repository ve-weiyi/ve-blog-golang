package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ ApiModel = (*defaultApiModel)(nil)

type (
	// 接口定义
	ApiModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ApiModel)
		// 插入
		Insert(ctx context.Context, in *Api) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Api) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *Api) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *Api) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Api, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Api, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Api, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Api, err error)
		// add extra method in here
		FindOneByPathMethodName(ctx context.Context, path string, method string, name string) (out *Api, err error)
	}

	// 表字段定义
	Api struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // 主键id
		ParentId  int64     `json:"parent_id" gorm:"column:parent_id" `   // 分组id
		Name      string    `json:"name" gorm:"column:name" `             // api名称
		Path      string    `json:"path" gorm:"column:path" `             // api路径
		Method    string    `json:"method" gorm:"column:method" `         // api请求方法
		Traceable int64     `json:"traceable" gorm:"column:traceable" `   // 是否追溯操作记录 0需要，1是
		Status    int64     `json:"status" gorm:"column:status" `         // 状态 1开，2关
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultApiModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewApiModel(db *gorm.DB, cache *redis.Client) ApiModel {
	return &defaultApiModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`api`",
	}
}

// 切换事务操作
func (m *defaultApiModel) WithTransaction(tx *gorm.DB) (out ApiModel) {
	return NewApiModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultApiModel) Insert(ctx context.Context, in *Api) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultApiModel) InsertBatch(ctx context.Context, in ...*Api) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultApiModel) Update(ctx context.Context, in *Api) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultApiModel) UpdateNotEmpty(ctx context.Context, in *Api) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultApiModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Api{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultApiModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Api{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultApiModel) FindOne(ctx context.Context, id int64) (out *Api, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultApiModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Api, err error) {
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
func (m *defaultApiModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Api{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultApiModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Api, err error) {
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
func (m *defaultApiModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Api, err error) {
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
func (m *defaultApiModel) FindOneByPathMethodName(ctx context.Context, path string, method string, name string) (out *Api, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`path` = ? and `method` = ? and `name` = ?", path, method, name).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
