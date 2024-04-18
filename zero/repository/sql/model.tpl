package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableName{{.UpperStartCamelName}} = "{{.SnakeName}}"

type (
	// 接口定义
	{{.UpperStartCamelName}}Model interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out {{.UpperStartCamelName}}Model)
		// 增删改查
		Create(ctx context.Context, in *{{.UpperStartCamelName}}) (out *{{.UpperStartCamelName}}, err error)
		Update(ctx context.Context, in *{{.UpperStartCamelName}}) (out *{{.UpperStartCamelName}}, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *{{.UpperStartCamelName}}, err error)
		// 批量操作
		InsertBatch(ctx context.Context, in ...*{{.UpperStartCamelName}}) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error)

	}

	// 接口实现
	default{{.UpperStartCamelName}}Model struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
        tableName 	string
	}

	// 表字段定义
	{{.UpperStartCamelName}} struct {
	{{range .Fields}}
        {{.Name}} {{.Type}} `{{.Tags}}` {{if .ColumnComment}}// {{.ColumnComment}}{{end}}
	{{- end}}
	}
)

func New{{.UpperStartCamelName}}Model(db *gorm.DB, cache *redis.Client) {{.UpperStartCamelName}}Model {
	return &default{{.UpperStartCamelName}}Model{
		DbEngin:    db,
		CacheEngin: cache,
		tableName: TableName{{.UpperStartCamelName}},
	}
}

// 切换事务操作
func (m *default{{.UpperStartCamelName}}Model) WithTransaction(tx *gorm.DB) (out {{.UpperStartCamelName}}Model) {
	return New{{.UpperStartCamelName}}Model(tx, m.CacheEngin)
}

// 创建{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) Create(ctx context.Context, in *{{.UpperStartCamelName}}) (out *{{.UpperStartCamelName}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) Update(ctx context.Context, in *{{.UpperStartCamelName}}) (out *{{.UpperStartCamelName}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&{{.UpperStartCamelName}}{})
	return query.RowsAffected, query.Error
}

// 查询{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) First(ctx context.Context, conditions string, args ...interface{}) (out *{{.UpperStartCamelName}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new({{.UpperStartCamelName}})
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) InsertBatch(ctx context.Context, in ...*{{.UpperStartCamelName}}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&{{.UpperStartCamelName}}{})
	return query.RowsAffected, query.Error
}

// 查询{{.UpperStartCamelName}}总数
func (m *default{{.UpperStartCamelName}}Model) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&{{.UpperStartCamelName}}{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询{{.UpperStartCamelName}}列表
func (m *default{{.UpperStartCamelName}}Model) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*{{.UpperStartCamelName}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

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

// 分页查询{{.UpperStartCamelName}}记录
func (m *default{{.UpperStartCamelName}}Model) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error) {
	// 创建db
	db := m.DbEngin.WithContext(ctx).Table(m.tableName)

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
