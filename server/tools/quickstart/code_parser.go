package quickstart

import (
	"strings"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/gorm_parser/field"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/gorm_parser/model"
)

// AutoCodeModel 初始版本自动化代码工具
type AutoCodeModel struct {
	DbName              string // 数据库名
	TableName           string // 表名 				auto_code
	UpperStartCamelName string // Struct名称 		AutoCode 大写驼峰命名
	LowerStartCamelName string // Struct变量名 		autoCode 小写驼峰命名
	SnakeName           string // StructJson名		auto_code api路径前缀
	CommentName         string // Struct中文名称 	「代码」	创建api的描述和注释
	IndexFields         map[string][]*field.Field
	Fields              []*field.Field
	ImportPkgPaths      []string
}

type TableParser struct {
	//Config
	DbEngin *gorm.DB
}

func NewTableParser(config Config) *TableParser {
	return &TableParser{
		//Config: config,
		DbEngin: config.DbEngin,
	}
}

func (t *TableParser) ParseModelFromSchema() []*AutoCodeModel {
	var models []*AutoCodeModel
	//dbName := t.DbEngin.Migrator().CurrentDatabase()
	tables, err := t.DbEngin.Migrator().GetTables()
	if err != nil {
		return nil
	}
	for _, table := range tables {
		m := t.ParseModelFromTable(table)
		if err != nil {
			return nil
		}
		models = append(models, m)
	}
	return models
}

func (t *TableParser) ParseModelFromTable(tableName string) *AutoCodeModel {
	dbName := t.DbEngin.Migrator().CurrentDatabase()
	tableInfos, err := t.DbEngin.Migrator().TableType(tableName)
	tableComment, _ := tableInfos.Comment()

	table, err := model.GetTable(t.DbEngin, tableName)
	if err != nil {
		return nil
	}

	out := &AutoCodeModel{
		DbName:              dbName,
		TableName:           tableName,
		UpperStartCamelName: jsonconv.Case2Camel(tableName),
		LowerStartCamelName: jsonconv.Case2CamelLowerStart(tableName),
		SnakeName:           jsonconv.Camel2Case(tableName),
		CommentName:         tableComment,
		Fields:              t.ConvertField(table.Columns),
		ImportPkgPaths:      []string{
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/controller",
			//"github.com/ve-weiyi/ve-blog-golang/server/api/common/model/response",
		},
	}

	return out
}

func (t *TableParser) ConvertField(columns []*model.Column) []*field.Field {
	var out []*field.Field
	for _, c := range columns {
		comment, _ := c.Comment()

		f := &field.Field{
			Name:             jsonconv.Case2Camel(c.Name()),
			Type:             c.FiledType(false, false, false),
			ColumnName:       c.Name(),
			ColumnComment:    comment,
			MultilineComment: strings.Contains(comment, "\n"),
			Tag:              map[string]string{field.TagKeyJson: jsonconv.Camel2Case(c.Name())},
			GORMTag:          c.BuildGormTag(),
		}

		out = append(out, f)
	}

	return out
}
