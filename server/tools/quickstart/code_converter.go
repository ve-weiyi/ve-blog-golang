package quickstart

import (
	"fmt"

	inject2 "github.com/ve-weiyi/ve-blog-golang/kit/tools/inject"
	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/tmpl"
)

type TableConverter struct {
	Config
	//OutPath   string                                   // 输出路径
	//OutFileNS func(tableName string) (fileName string) // 输出文件名称
	//
	//ReplaceMode int //是否替换文件 0:创建或替换 1:只创建 2:只替换
}

func NewTableConverter(config Config) *TableConverter {
	return &TableConverter{
		Config: config,
	}
}

func (t *TableConverter) GenerateInjectMetas(models ...*AutoCodeModel) []*inject2.AstInjectMeta {
	var injectMetas []*inject2.AstInjectMeta

	for _, data := range models {
		temporaryRoot := t.OutPath

		injectMetas = append(injectMetas, &inject2.AstInjectMeta{
			Key:      tmpl.KeyRepository,
			FilePath: fmt.Sprintf("%v/repository/repository.go", temporaryRoot),
			StructMetas: []*inject2.StructMeta{
				inject2.NewStructMete("AppRepository", fmt.Sprintf(`%vRepository *logic.%vRepository //%v`, data.UpperStartCamelName, data.UpperStartCamelName, data.CommentName)),
			},
			FuncMetas: []*inject2.FuncMeta{
				inject2.NewFuncMete("NewRepository", fmt.Sprintf(`return &AppRepository{
			%vRepository: logic.New%vRepository(svcCtx),
			}`, data.UpperStartCamelName, data.UpperStartCamelName)),
			},
		})

		injectMetas = append(injectMetas, &inject2.AstInjectMeta{
			Key:      tmpl.KeyService,
			FilePath: fmt.Sprintf("%v/service.go", temporaryRoot),
			StructMetas: []*inject2.StructMeta{
				inject2.NewStructMete("AppService", fmt.Sprintf(`%vService *logic.%vService //%v`, data.UpperStartCamelName, data.UpperStartCamelName, data.CommentName)),
			},
			FuncMetas: []*inject2.FuncMeta{
				inject2.NewFuncMete("NewService", fmt.Sprintf(`return &AppService{
			%vService: logic.New%vService(svcCtx),
			}`, data.UpperStartCamelName, data.UpperStartCamelName)),
			},
		})

		injectMetas = append(injectMetas, &inject2.AstInjectMeta{
			Key:      tmpl.KeyController,
			FilePath: fmt.Sprintf("%v/controller/controller.go", temporaryRoot),
			StructMetas: []*inject2.StructMeta{
				inject2.NewStructMete("AppController", fmt.Sprintf(`%vController *logic.%vController //%v`, data.UpperStartCamelName, data.UpperStartCamelName, data.CommentName)),
			},
			FuncMetas: []*inject2.FuncMeta{
				inject2.NewFuncMete("NewController", fmt.Sprintf(`return &AppController{
			%vController: logic.New%vController(svcCtx),
			}`, data.UpperStartCamelName, data.UpperStartCamelName)),
			},
		})

		injectMetas = append(injectMetas, &inject2.AstInjectMeta{
			Key:      tmpl.KeyRouter,
			FilePath: fmt.Sprintf("%v/router/router.go", temporaryRoot),
			StructMetas: []*inject2.StructMeta{
				inject2.NewStructMete("AppRouter", fmt.Sprintf(`%vRouter *logic.%vRouter //%v`, data.UpperStartCamelName, data.UpperStartCamelName, data.CommentName)),
			},
			FuncMetas: []*inject2.FuncMeta{
				inject2.NewFuncMete("NewRouter", fmt.Sprintf(`return &AppRouter{
			%vRouter: logic.New%vRouter(svcCtx),
			}`, data.UpperStartCamelName, data.UpperStartCamelName)),
			},
		})

		injectMetas = append(injectMetas, &inject2.AstInjectMeta{
			Key:      tmpl.KeyRouter,
			FilePath: fmt.Sprintf("%v/router/logic/register.rt.go", temporaryRoot),
			DeclMetas: []*inject2.DeclMeta{inject2.NewDeclMeta(fmt.Sprintf(`
	// 初始化 %s 路由信息
	// publicRouter 公开路由，不登录就可以访问
	// loginRouter  登录路由，登录后才可以访问
	func (s *%sRouter) Init%sRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
		s.Init%sBasicRouter(publicRouter, loginRouter)
	}
`, data.UpperStartCamelName, data.UpperStartCamelName, data.UpperStartCamelName, data.UpperStartCamelName))},
		})
	}
	return injectMetas
}

func (t *TableConverter) GenerateInventMetas(models ...*AutoCodeModel) []*invent.TemplateMeta {
	var metas []*invent.TemplateMeta

	for _, data := range models {
		temporaryRoot := t.OutPath
		fileName := t.OutFileNS(data.TableName)
		mode := t.ReplaceMode
		metas = append(metas, &invent.TemplateMeta{
			Key:            tmpl.KeyModel,
			TemplateString: tmpl.Model,
			CodeOutPath:    fmt.Sprintf("%v/model/entity/%s.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			Key:            tmpl.KeyRepository,
			TemplateString: tmpl.Repository,
			CodeOutPath:    fmt.Sprintf("%v/repository/logic/%s.rp.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			Key:            tmpl.KeyService,
			TemplateString: tmpl.Service,
			CodeOutPath:    fmt.Sprintf("%v/service/logic/%s.sv.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			Key:            tmpl.KeyController,
			TemplateString: tmpl.Controller,
			CodeOutPath:    fmt.Sprintf("%v/controller/logic/%s.ctl.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			Key:            tmpl.KeyRouter,
			TemplateString: tmpl.Router,
			CodeOutPath:    fmt.Sprintf("%v/router/logic/%s.rt.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		//metas = append(metas, &invent.TemplateMeta{
		//	Key:            tmpl.KeyApi,
		//	TemplateString: tmpl.Api,
		//	CodeOutPath:    fmt.Sprintf("%v/api/%s.ts", temporaryRoot, fileName),
		//	Data:           data,
		//	Mode:           mode,
		//})
	}

	return metas
}

func (t *TableConverter) GenerateCommonInventMetas(models ...*AutoCodeModel) []*invent.TemplateMeta {
	var metas []*invent.TemplateMeta

	for _, data := range models {
		temporaryRoot := t.OutPath
		fileName := t.OutFileNS(data.TableName)
		mode := t.ReplaceMode

		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.CommonRepository,
			CodeOutPath:    fmt.Sprintf("%v/repository/logic/%s.rp.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})

		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.CommonService,
			CodeOutPath:    fmt.Sprintf("%v/service/logic/%s.sv.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})

		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.CommonController,
			CodeOutPath:    fmt.Sprintf("%v/controller/logic/%s.ctl.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.CommonRouter,
			CodeOutPath:    fmt.Sprintf("%v/router/logic/%s.rt.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
	}
	return metas
}

func (t *TableConverter) GeneratePkgMetas(models ...*AutoCodeModel) ([]*invent.TemplateMeta, []*inject2.AstInjectMeta) {
	var metas []*invent.TemplateMeta

	for _, data := range models {
		temporaryRoot := t.OutPath
		fileName := "context"
		mode := t.ReplaceMode

		//context 是下层引用
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.RouterContext,
			CodeOutPath:    fmt.Sprintf("%v/router/svc/%s.rt.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.ControllerContext,
			CodeOutPath:    fmt.Sprintf("%v/controller/svc/%s.ctl.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.ServiceContext,
			CodeOutPath:    fmt.Sprintf("%v/service/svc/%s.sv.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.RepositoryContext,
			CodeOutPath:    fmt.Sprintf("%v/repository/svc/%s.rp.go", temporaryRoot, fileName),
			Data:           data,
			Mode:           mode,
		})

		//入口文件
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.AppRouter,
			CodeOutPath:    fmt.Sprintf("%v/router/router.go", temporaryRoot),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.AppController,
			CodeOutPath:    fmt.Sprintf("%v/controller/controller.go", temporaryRoot),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.AppService,
			CodeOutPath:    fmt.Sprintf("%v/service.go", temporaryRoot),
			Data:           data,
			Mode:           mode,
		})
		metas = append(metas, &invent.TemplateMeta{
			TemplateString: tmpl.AppRepository,
			CodeOutPath:    fmt.Sprintf("%v/repository/repository.go", temporaryRoot),
			Data:           data,
			Mode:           mode,
		})
	}

	return metas, nil
}
