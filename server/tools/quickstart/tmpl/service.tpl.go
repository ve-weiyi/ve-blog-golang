package tmpl

const AppService = `
package service

import (

)

type AppService struct {
	svcCtx *svc.ServiceContext //持有的repository层引用
}

func NewService(svcCtx *svc.ServiceContext) *AppService {
	return &AppService{
		svcCtx: svcCtx,
	}
}
`

const ServiceContext = `
package svc

import (

)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	*repository.AppRepository
	Config *config.Config
	Log    *glog.Glogger
	//下面是引用的repository
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	ctx := svc.NewRepositoryContext(cfg)
	repo := repository.NewRepository(ctx)
	if repo == nil {
		panic("repository cannot be null")
	}

	return &ServiceContext{
		AppRepository: repo,
		Config:        cfg,
		Log:           global.LOG,
	}
}

`

const Service = `
package logic

import (
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Service struct {
	svcCtx *svc.ServiceContext
}

func New{{.StructName}}Service(svcCtx *svc.ServiceContext) *{{.StructName}}Service {
	return &{{.StructName}}Service{
		svcCtx: svcCtx,
	}
}

// 创建{{.StructName}}记录
func (s *{{.StructName}}Service) Create{{.StructName}}(reqCtx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	return s.svcCtx.{{.StructName}}Repository.Create(reqCtx, {{.ValueName}})
}

// 更新{{.StructName}}记录
func (s *{{.StructName}}Service) Update{{.StructName}}(reqCtx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	return s.svcCtx.{{.StructName}}Repository.Update(reqCtx, {{.ValueName}})
}

// 删除{{.StructName}}记录
func (s *{{.StructName}}Service) Delete{{.StructName}}(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.{{.StructName}}Repository.Delete(reqCtx, "id = ?", id)
}

// 查询{{.StructName}}记录
func (s *{{.StructName}}Service) Find{{.StructName}}(reqCtx *request.Context, id int) (data *entity.{{.StructName}}, err error) {
	return s.svcCtx.{{.StructName}}Repository.First(reqCtx, "id = ?", id)
}

// 批量删除{{.StructName}}记录
func (s *{{.StructName}}Service) Delete{{.StructName}}ByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.{{.StructName}}Repository.Delete(reqCtx, "id in (?)", ids)
}

// 分页获取{{.StructName}}记录
func (s *{{.StructName}}Service) Find{{.StructName}}List(reqCtx *request.Context, page *request.PageQuery) (list []*entity.{{.StructName}}, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = s.svcCtx.{{.StructName}}Repository.FindList(reqCtx, page.Page, page.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.{{.StructName}}Repository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
`

const CommonService = `
package logic

import (

)

type {{.StructName}}Service struct {
	svcCtx *svc.ServiceContext
}

func New{{.StructName}}Service(svcCtx *svc.ServiceContext) *{{.StructName}}Service {
	return &{{.StructName}}Service{
		svcCtx: svcCtx,
	}
}

`
