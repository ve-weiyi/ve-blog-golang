package tmpl

const AppService = `
package service

import (
	"{{.SvcPackage }}"
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
	Config *config.Config
	Log    *glog.Glogger
	//下面是引用的repository
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: cfg,
		Log:    global.LOG,

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
	return s.svcCtx.{{.StructName}}Repository.Create{{.StructName}}(reqCtx, {{.ValueName}})
}

// 删除{{.StructName}}记录
func (s *{{.StructName}}Service) Delete{{.StructName}}(reqCtx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (rows int64, err error) {
	return s.svcCtx.{{.StructName}}Repository.Delete{{.StructName}}(reqCtx, {{.ValueName}})
}

// 更新{{.StructName}}记录
func (s *{{.StructName}}Service) Update{{.StructName}}(reqCtx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	return s.svcCtx.{{.StructName}}Repository.Update{{.StructName}}(reqCtx, {{.ValueName}})
}

// 查询{{.StructName}}记录
func (s *{{.StructName}}Service) Get{{.StructName}}(reqCtx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	return s.svcCtx.{{.StructName}}Repository.Get{{.StructName}}(reqCtx, {{.ValueName}}.ID)
}

// 批量删除{{.StructName}}记录
func (s *{{.StructName}}Service) Delete{{.StructName}}ByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.{{.StructName}}Repository.Delete{{.StructName}}ByIds(reqCtx, ids)
}

// 分页获取{{.StructName}}记录
func (s *{{.StructName}}Service) Find{{.StructName}}List(reqCtx *request.Context, page *request.PageInfo) (list []*entity.{{.StructName}}, total int64, err error) {
	return s.svcCtx.{{.StructName}}Repository.Find{{.StructName}}List(reqCtx, page)
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
