package easycode

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/tmpl"
)

func (g *Generator) InitPackage(tableName string) {
	data := '1'
	temporaryRoot := g.cfg.OutPath
	fileName := "context"

	replace := false
	//context 是下层引用
	routerContext := &plate.PlateMeta{
		TemplateString: tmpl.RouterContext,
		AutoCodePath:   fmt.Sprintf("%v/router/svc/%s.rt.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        replace,
	}
	controllerContext := &plate.PlateMeta{
		TemplateString: tmpl.ControllerContext,
		AutoCodePath:   fmt.Sprintf("%v/controller/svc/%s.ctl.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        replace,
	}
	serviceContext := &plate.PlateMeta{
		TemplateString: tmpl.ServiceContext,
		AutoCodePath:   fmt.Sprintf("%v/service/svc/%s.sv.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        replace,
	}
	repositoryContext := &plate.PlateMeta{
		TemplateString: tmpl.RepositoryContext,
		AutoCodePath:   fmt.Sprintf("%v/repository/svc/%s.rp.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        replace,
	}

	metas := []*plate.PlateMeta{
		/** server start */
		routerContext,
		controllerContext,
		serviceContext,
		repositoryContext,
	}

	for _, item := range metas {
		err := item.CreateTempFile()
		if err != nil {
			g.Logger.Printf("%v:%v", item.AutoCodePath, err)
		}
	}

	routerCollector := &plate.PlateMeta{
		TemplateString: tmpl.AppRouter,
		AutoCodePath:   fmt.Sprintf("%v/router/router.go", temporaryRoot),
		Data:           data,
		Replace:        replace,
	}
	controllerCollector := &plate.PlateMeta{
		TemplateString: tmpl.AppController,
		AutoCodePath:   fmt.Sprintf("%v/controller/controller.go", temporaryRoot),
		Data:           data,
		Replace:        replace,
	}
	serviceCollector := &plate.PlateMeta{
		TemplateString: tmpl.AppService,
		AutoCodePath:   fmt.Sprintf("%v/service/service.go", temporaryRoot),
		Data:           data,
		Replace:        replace,
	}
	repositoryCollector := &plate.PlateMeta{
		TemplateString: tmpl.AppRepository,
		AutoCodePath:   fmt.Sprintf("%v/repository/repository.go", temporaryRoot),
		Data:           data,
		Replace:        replace,
	}

	metasPkg := []*plate.PlateMeta{
		/** server start */
		routerCollector,
		controllerCollector,
		serviceCollector,
		repositoryCollector,
	}
	for _, item := range metasPkg {
		err := item.CreateTempFile()
		if err != nil {
			g.Logger.Printf("%v:%v", item.AutoCodePath, err)
		}
	}
}
