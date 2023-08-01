package easycode

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/inject"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate/field"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate/provider"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/tmpl"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/dbdriver"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

type Generator struct {
	cfg Config
	*log.Logger

	plateMetas  []*plate.PlateMeta
	InjectMetas []*inject.AstInjectMeta

	plateData []*plate.AutoCodeStructData

	customPlateMetas []*plate.PlateMeta
}

// UseDB set db connection
func (g *Generator) UseDB(db *gorm.DB) {
	if db != nil {
		g.cfg.db = db
	}
}

func NewGenerator(cfg Config) *Generator {
	if err := cfg.Revise(); err != nil {
		panic(fmt.Errorf("create generator fail: %w", err))
	}
	return &Generator{
		cfg:    cfg,
		Logger: log.New(os.Stderr, "", log.LstdFlags|log.Llongfile),
	}
}

// Execute generate code to output path
func (g *Generator) Execute() {
	g.Println("Start generating code.")

	if err := g.generateModelFile(); err != nil {
		g.Printf("generate model struct fail: %s", err)
		panic("generate model struct fail")
	}

	g.Println("Generate code done.")
}
func (g *Generator) RollBack() {
	g.Println("Start rollback code.")

	if err := g.rollback(); err != nil {
		g.Printf("rollback model struct fail: %s", err)
		panic("rollback model struct fail")
	}

	g.Println("RollBack code done.")
}

func (g *Generator) GenFieldConfig() *field.FieldConfig {
	return &field.FieldConfig{
		DataTypeMap: g.cfg.dataTypeMap,

		FieldSignable:     g.cfg.FieldSignable,
		FieldNullable:     g.cfg.FieldNullable,
		FieldCoverable:    g.cfg.FieldCoverable,
		FieldWithIndexTag: g.cfg.FieldWithIndexTag,
		FieldWithTypeTag:  g.cfg.FieldWithTypeTag,

		FieldNameNS:  g.cfg.FieldNameNS,
		FieldJsonNS:  g.cfg.FieldJsonNS,
		FieldValueNS: g.cfg.FieldValueNS,
	}
}

func (g *Generator) ApplyMetas(plates []*plate.PlateMeta, injects []*inject.AstInjectMeta) {
	g.plateMetas = append(g.plateMetas, plates...)

	g.InjectMetas = append(g.InjectMetas, injects...)
}

// 创建数据库中所有表
func (g *Generator) GenerateMetasFromSchema() ([]*plate.PlateMeta, []*inject.AstInjectMeta) {
	db := g.cfg.db
	//去拿到表注释
	mysqlDriver := dbdriver.MysqlDriver{DB: db}
	tbs, err := mysqlDriver.GetTables(db.Migrator().CurrentDatabase())
	if err != nil {
		return nil, nil
	}

	var pl []*plate.PlateMeta
	var in []*inject.AstInjectMeta
	for _, tb := range tbs {
		plates, injects := g.GenerateMetasFromTable(tb.TableName, tb.TableComment)
		pl = append(pl, plates...)
		in = append(in, injects...)
	}
	return pl, in
}

// 创建一个表
func (g *Generator) GenerateMetasFromTable(tableName, tableComment string) ([]*plate.PlateMeta, []*inject.AstInjectMeta) {
	db := g.cfg.db
	columns, err := provider.GetTableColumns(db, tableName)
	if err != nil {
		return nil, nil
	}

	cfg := g.GenFieldConfig()
	var fields []*provider.Field
	for _, column := range columns {
		field := column.ToField(cfg)
		fields = append(fields, field)
	}

	if tableComment == "" {
		tableComment = tableName
	}

	return g.GenerateMetasFromModel(tableName, tableComment, fields)
}

// 使用data进行创建
func (g *Generator) GenerateMetasFromModel(tableName, tableComment string, fields []*provider.Field) ([]*plate.PlateMeta, []*inject.AstInjectMeta) {
	db := g.cfg.db
	dbName := db.Migrator().CurrentDatabase()

	data := &plate.AutoCodeStructData{
		Package:        jsonconv.Case2CamelNotFirst(dbName),
		TableName:      tableName,
		StructName:     jsonconv.Case2Camel(tableName),
		ValueName:      jsonconv.Case2CamelNotFirst(tableName),
		JsonName:       jsonconv.Camel2Case(tableName),
		StructComment:  tableComment,
		Fields:         fields,
		ImportPkgPaths: []string{
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/controller",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/model/response",
		},
	}
	g.plateData = append(g.plateData, data)

	temporaryRoot := g.cfg.OutPath
	fileName := g.cfg.OutFileNS(tableName)

	metaModel := &plate.PlateMeta{
		Key:            tmpl.KeyModel,
		TemplateString: tmpl.Model,
		AutoCodePath:   fmt.Sprintf("%v/model/entity/%s.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}
	metaRepository := &plate.PlateMeta{
		Key:            tmpl.KeyRepository,
		TemplateString: tmpl.Repository,
		AutoCodePath:   fmt.Sprintf("%v/repository/logic/%s.rp.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}
	metaService := &plate.PlateMeta{
		Key:            tmpl.KeyService,
		TemplateString: tmpl.Service,
		AutoCodePath:   fmt.Sprintf("%v/service/logic/%s.sv.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}
	metaController := &plate.PlateMeta{
		Key:            tmpl.KeyController,
		TemplateString: tmpl.Controller,
		AutoCodePath:   fmt.Sprintf("%v/controller/logic/%s.ctl.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}
	metaRouter := &plate.PlateMeta{
		Key:            tmpl.KeyRouter,
		TemplateString: tmpl.Router,
		AutoCodePath:   fmt.Sprintf("%v/router/logic/%s.rt.go", temporaryRoot, fileName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}
	metaApi := &plate.PlateMeta{
		Key:            tmpl.KeyApi,
		TemplateString: tmpl.Api,
		AutoCodePath:   fmt.Sprintf("%v/api/%s.ts", temporaryRoot, tableName),
		Data:           data,
		Replace:        g.cfg.Replace,
	}

	metas := []*plate.PlateMeta{
		/** server start */
		metaModel,
		metaRepository,
		metaService,
		metaController,
		metaRouter,
		metaApi,
	}

	var injectMetas []*inject.AstInjectMeta

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		Key:      tmpl.KeyRepository,
		FilePath: fmt.Sprintf("%v/repository/repository.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppRepository", fmt.Sprintf(`%vRepository *logic.%vRepository //%v`, data.StructName, data.StructName, data.StructComment)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewRepository", fmt.Sprintf(`return &AppRepository{
			%vRepository: logic.New%vRepository(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		Key:      tmpl.KeyService,
		FilePath: fmt.Sprintf("%v/service/service.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppService", fmt.Sprintf(`%vService *logic.%vService //%v`, data.StructName, data.StructName, data.StructComment)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewService", fmt.Sprintf(`return &AppService{
			%vService: logic.New%vService(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		Key:      tmpl.KeyController,
		FilePath: fmt.Sprintf("%v/controller/controller.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppController", fmt.Sprintf(`%vController *logic.%vController //%v`, data.StructName, data.StructName, data.StructComment)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewController", fmt.Sprintf(`return &AppController{
			%vController: logic.New%vController(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		Key:      tmpl.KeyRouter,
		FilePath: fmt.Sprintf("%v/router/router.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppRouter", fmt.Sprintf(`%vRouter *logic.%vRouter //%v`, data.StructName, data.StructName, data.StructComment)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewRouter", fmt.Sprintf(`return &AppRouter{
			%vRouter: logic.New%vRouter(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		Key:      tmpl.KeyRouter,
		FilePath: fmt.Sprintf("%v/router/logic/register.rt.go", temporaryRoot),
		DeclMeta: []*inject.DeclMeta{inject.NewDeclMeta(fmt.Sprintf(`
	// 初始化 %s 路由信息
	// publicRouter 公开路由，不登录就可以访问
	// loginRouter  登录路由，登录后才可以访问
	func (s *%sRouter) Init%sRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
		s.Init%sGenRouter(publicRouter, loginRouter)
	}
`, data.StructName, data.StructName, data.StructName, data.StructName))},
	})
	return metas, injectMetas
}

func (g *Generator) GetTemplateDataList() []*plate.AutoCodeStructData {
	return g.plateData
}

func (g *Generator) GenerateCommonFile(tableName string, tableComment string) error {
	db := g.cfg.db
	dbName := db.Migrator().CurrentDatabase()

	data := &plate.AutoCodeStructData{
		Package:        jsonconv.Case2CamelNotFirst(dbName),
		StructName:     jsonconv.Case2Camel(tableName),
		ValueName:      jsonconv.Case2CamelNotFirst(tableName),
		JsonName:       jsonconv.Camel2Case(tableName),
		StructComment:  tableComment,
		Fields:         nil,
		ImportPkgPaths: []string{
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/controller",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/model/response",
		},
	}

	temporaryRoot := g.cfg.OutPath
	if g.cfg.GenerateCommon {
		fileName := fmt.Sprintf("%v", tableName)

		//metaCommonDao := &plate.PlateMeta{
		//	TemplateString: tmpl.CommonDao,
		//	AutoCodePath:   fmt.Sprintf("%v/model/dao/%s.sv.go", temporaryRoot, fileName),
		//	EncodeData:           data,
		//	Replace:        g.cfg.ReplaceCommon,
		//}

		metaCommonRepository := &plate.PlateMeta{
			TemplateString: tmpl.CommonRepository,
			AutoCodePath:   fmt.Sprintf("%v/repository/logic/%s.rp.go", temporaryRoot, fileName),
			Data:           data,
			Replace:        g.cfg.ReplaceCommon,
		}

		metaCommonService := &plate.PlateMeta{
			TemplateString: tmpl.CommonService,
			AutoCodePath:   fmt.Sprintf("%v/service/logic/%s.sv.go", temporaryRoot, fileName),
			Data:           data,
			Replace:        g.cfg.ReplaceCommon,
		}

		metaCommonController := &plate.PlateMeta{
			TemplateString: tmpl.CommonController,
			AutoCodePath:   fmt.Sprintf("%v/controller/logic/%s.ctl.go", temporaryRoot, fileName),
			Data:           data,
			Replace:        g.cfg.ReplaceCommon,
		}
		metaCommonRouter := &plate.PlateMeta{
			TemplateString: tmpl.CommonRouter,
			AutoCodePath:   fmt.Sprintf("%v/router/logic/%s.rt.go", temporaryRoot, fileName),
			Data:           data,
			Replace:        g.cfg.ReplaceCommon,
		}

		g.customPlateMetas = append(g.customPlateMetas,
			//metaCommonDao,
			metaCommonRepository,
			metaCommonService,
			metaCommonController,
			metaCommonRouter)
	}

	for _, item := range g.customPlateMetas {
		err := item.CreateTempFile()
		if err != nil {
			g.Logger.Printf("%v:%v", item.AutoCodePath, err)
		}
	}

	return nil
}

func (g *Generator) generateModelFile() error {

	for _, item := range g.plateMetas {
		if _, ok := g.cfg.GenerateMap[item.Key]; !ok {
			continue
		}
		err := item.CreateTempFile()
		if err != nil {
			g.Logger.Printf("%v:%v", item.AutoCodePath, err)
		}
	}

	for _, item := range g.InjectMetas {
		if _, ok := g.cfg.GenerateMap[item.Key]; !ok {
			continue
		}
		err := item.Inject()
		if err != nil {
			g.Logger.Println(err)
		}
	}
	return nil
}

func (g *Generator) rollback() error {
	for _, item := range g.plateMetas {
		if _, ok := g.cfg.GenerateMap[item.Key]; !ok {
			continue
		}
		err := item.RollBack()
		if err != nil {
			return err
		}
	}

	for _, item := range g.InjectMetas {
		if _, ok := g.cfg.GenerateMap[item.Key]; !ok {
			continue
		}
		err := item.RollBack()
		if err != nil {
			return err
		}
	}

	return nil
}
