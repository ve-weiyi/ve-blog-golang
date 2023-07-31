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
func (s *{{.StructName}}Repository) Create{{.StructName}}(ctx context.Context, {{.ValueName}} *entity.{{.StructName}}) (out *entity.{{.StructName}}, err error) {
	db:= s.DbEngin
	err = db.Create(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 更新{{.StructName}}记录
func (s *{{.StructName}}Repository) Update{{.StructName}}(ctx context.Context, {{.ValueName}} *entity.{{.StructName}}) (out *entity.{{.StructName}}, err error) {
	db:= s.DbEngin	
	err = db.Save(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 删除{{.StructName}}记录
func (s *{{.StructName}}Repository) Delete{{.StructName}}(ctx context.Context, id int) (rows int64, err error) {
	db:= s.DbEngin	
	query := db.Delete(&entity.{{.StructName}}{}, "id = ?", id)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询{{.StructName}}记录
func (s *{{.StructName}}Repository) Find{{.StructName}}(ctx context.Context, id int) (out *entity.{{.StructName}}, err error) {
	db:= s.DbEngin	
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除{{.StructName}}记录
func (s *{{.StructName}}Repository) Delete{{.StructName}}ByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db:= s.DbEngin	
	query := db.Delete(&entity.{{.StructName}}{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询{{.StructName}}记录
func (s *{{.StructName}}Repository) Find{{.StructName}}List(ctx context.Context, page *request.PageQuery) (list []*entity.{{.StructName}}, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
		db = db.Order(page.OrderClause())
	}

	// 查询总数,要在使用limit之前
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
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
