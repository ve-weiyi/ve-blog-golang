package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ {{.UpperStartCamelName}}Model = (*default{{.UpperStartCamelName}}Model)(nil)

type (
	// 接口定义
	{{.UpperStartCamelName}}Model interface {
		TableName() string
		// 在事务中操作
		WithTransaction(tx *gorm.DB) (out {{.UpperStartCamelName}}Model)
		// 插入
		Insert(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*{{.UpperStartCamelName}}) (rows int64, err error)
        // 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
        // 更新
        Update(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error)
        Save(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *{{.UpperStartCamelName}}, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *{{.UpperStartCamelName}}, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error)
	    // add extra method in here
        {{- range $key, $value := .UniqueFields}}
            FindOneBy{{ funcFieldsKey $value}}(ctx context.Context, {{funcFieldsKeyVar $value}}) (out *{{$.UpperStartCamelName}},err error)
        {{- end}}
	}

	// 表字段定义
	{{.UpperStartCamelName}} struct {
	{{range .Fields}}
        {{.Name}} {{.Type}} `{{.Tags}}` {{if .Comment}}// {{.Comment}}{{end}}
	{{- end}}
	}

    // 接口实现
    default{{.UpperStartCamelName}}Model struct {
        DbEngin    *gorm.DB
        CacheEngin *redis.Client
        table       string
    }
)

func New{{.UpperStartCamelName}}Model(db *gorm.DB, cache *redis.Client) {{.UpperStartCamelName}}Model {
	return &default{{.UpperStartCamelName}}Model{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`{{.SnakeName}}`",
	}
}

func (m *default{{.UpperStartCamelName}}Model) TableName() string {
	return m.table
}

// 在事务中操作
func (m *default{{.UpperStartCamelName}}Model) WithTransaction(tx *gorm.DB) (out {{.UpperStartCamelName}}Model) {
	return New{{.UpperStartCamelName}}Model(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *default{{.UpperStartCamelName}}Model) Insert(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *default{{.UpperStartCamelName}}Model) InsertBatch(ctx context.Context, in ...*{{.UpperStartCamelName}}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *default{{.UpperStartCamelName}}Model) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&{{.UpperStartCamelName}}{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *default{{.UpperStartCamelName}}Model) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&{{.UpperStartCamelName}}{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 保存记录（更新零值）
func (m *default{{.UpperStartCamelName}}Model) Save(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *default{{.UpperStartCamelName}}Model) Update(ctx context.Context, in *{{.UpperStartCamelName}}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *default{{.UpperStartCamelName}}Model) FindOne(ctx context.Context, id int64) (out *{{.UpperStartCamelName}}, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *default{{.UpperStartCamelName}}Model) First(ctx context.Context, conditions string, args ...interface{}) (out *{{.UpperStartCamelName}}, err error) {
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
func (m *default{{.UpperStartCamelName}}Model) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 查询列表
func (m *default{{.UpperStartCamelName}}Model) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*{{.UpperStartCamelName}}, err error) {
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
func (m *default{{.UpperStartCamelName}}Model) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*{{.UpperStartCamelName}}, err error) {
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
{{- range $key, $value := .UniqueFields}}
func (m *default{{$.UpperStartCamelName}}Model) FindOneBy{{ funcFieldsKey $value}}(ctx context.Context, {{funcFieldsKeyVar $value}}) (out *{{$.UpperStartCamelName}},err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("{{funcFieldsKeyCond $value}}", {{funcFieldsKeyCondVar $value}}).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
{{- end}}
