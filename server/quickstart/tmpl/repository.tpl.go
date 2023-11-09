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
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 更新{{.StructName}}记录
func (s *{{.StructName}}Repository) Update{{.StructName}}(ctx context.Context, {{.ValueName}} *entity.{{.StructName}}) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 删除{{.StructName}}记录
func (s *{{.StructName}}Repository) Delete{{.StructName}}(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.{{.StructName}}{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询{{.StructName}}记录
func (s *{{.StructName}}Repository) Find{{.StructName}}(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询{{.StructName}}记录
func (s *{{.StructName}}Repository) Find{{.StructName}}List(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.{{.StructName}}, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sqlx.OrderClause(sorts))
	}

	// 如果有分页参数
	if page != nil && page.IsValid() {
		limit := page.Limit()
		offset := page.Offset()
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
func (s *{{.StructName}}Repository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.{{.StructName}}{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询{{.StructName}}记录——根据id
func (s *{{.StructName}}Repository) Find{{.StructName}}ById(ctx context.Context, id int) (out *entity.{{.StructName}}, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}


// 删除{{.StructName}}记录——根据id
func (s *{{.StructName}}Repository) Delete{{.StructName}}ById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.{{.StructName}}{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除{{.StructName}}记录——根据ids
func (s *{{.StructName}}Repository) Delete{{.StructName}}ByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.{{.StructName}}{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
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
