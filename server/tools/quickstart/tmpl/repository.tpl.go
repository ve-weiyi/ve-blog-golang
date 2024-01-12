package tmpl

const AppRepository = `
package repository

import (

)

//对应go-zero model层服务
type AppRepository struct {
	svcCtx *svc.RepositoryContext //持有的repository层引用
}

func NewRepository(svcCtx *svc.RepositoryContext) *AppRepository {
	return &AppRepository{
		svcCtx: svcCtx,
	}
}
`

const RepositoryContext = `
package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 注册需要用到的gorm、redis、model
type RepositoryContext struct {
	Config   *config.Config
	DbEngin  *gorm.DB
	DBList   map[string]*gorm.DB
	Cache    *redis.Client
	Log      *glog.Glogger
	//下面是一些Model
}

func NewRepositoryContext(cfg *config.Config) *RepositoryContext {
	return &RepositoryContext{
		Config:  cfg,
		DbEngin: global.DB,
		DBList:  global.DBList,
		Cache:   global.REDIS,
		Log:     global.LOG,
	}
}

`

const Repository = `
package logic

import (
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Repository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func New{{.StructName}}Repository(svcCtx *svc.RepositoryContext) *{{.StructName}}Repository {
	return &{{.StructName}}Repository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建{{.StructName}}记录
func (s *{{.StructName}}Repository) Create(ctx context.Context, item *entity.{{.StructName}}) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item, err
}

// 更新{{.StructName}}记录
func (s *{{.StructName}}Repository) Update(ctx context.Context, item *entity.{{.StructName}}) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&item).Error
	if err != nil {
		return nil, err
	}
	return item, err
}

// 删除{{.StructName}}记录
func (s *{{.StructName}}Repository) Delete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&entity.{{.StructName}}{})
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询{{.StructName}}记录
func (s *{{.StructName}}Repository) First(ctx context.Context, conditions string, args ...interface{}) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

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

func (s *{{.StructName}}Repository) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

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

// 分页查询{{.StructName}}记录
func (s *{{.StructName}}Repository) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*entity.{{.StructName}}, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

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

// 查询总数
func (s *{{.StructName}}Repository) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&entity.{{.StructName}}{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
`

const CommonRepository = `
package logic

import (

)

type {{.StructName}}Repository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func New{{.StructName}}Repository(svcCtx *svc.RepositoryContext) *{{.StructName}}Repository {
	return &{{.StructName}}Repository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}
`
