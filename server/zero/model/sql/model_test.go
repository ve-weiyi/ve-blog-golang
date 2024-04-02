package sql

import (
	_ "embed"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/invent/field"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/parser"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

//go:embed model.tpl
var modelTpl string

func TestParseSql(t *testing.T) {
	tables, err := parser.Parse(filepath.Join(global.GetRuntimeRoot(), "server/test.sql"), "go_zero", false)
	//tables, err := parser.Parse(filepath.Join(global.GetCurrentDir(), "t_user.sql"), "go_zero", false)
	assert.Nil(t, err)

	//log.Printf("table %+v", jsonconv.ObjectToJsonIndent(tables))

	for _, table := range tables {
		log.Printf("%+v", table.Name)
		log.Printf("%+v", jsonconv.ObjectToJsonIndent(table.UniqueIndex))

		metas := NewMetas(table)

		for _, meta := range metas {
			err = meta.CreateTempFile()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func NewMetas(table *parser.Table) []invent.TemplateMeta {
	var fs []*field.Field
	for _, e := range table.Fields {
		//log.Printf("%+v", jsonconv.ObjectToJsonIndent(e))

		fs = append(fs, convertField(e))
	}

	ifs := make(map[string][]*field.Field)
	for k, v := range table.UniqueIndex {
		var f []*field.Field
		for _, e := range v {
			f = append(f, convertField(e))
		}
		ifs[k] = f
	}

	var metas []invent.TemplateMeta

	metas = append(metas, invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fmt.Sprintf("../%v.model.go", table.Name.Source()),
		TemplateString: modelTpl,
		FunMap:         nil,
		Data: quickstart.AutoCodeModel{
			DbName:              table.Db.Source(),
			TableName:           table.Name.Source(),
			UpperStartCamelName: jsonconv.Case2Camel(table.Name.Source()),
			LowerStartCamelName: jsonconv.Case2CamelLowerStart(table.Name.Source()),
			SnakeName:           jsonconv.Camel2Case(table.Name.Source()),
			CommentName:         table.Name.Source(),
			IndexFields:         ifs,
			Fields:              fs,
			ImportPkgPaths:      nil,
		},
	})

	return metas
}

func convertField(e *parser.Field) *field.Field {
	return &field.Field{
		Name: jsonconv.Case2Camel(e.Name.Source()),
		Type: func() string {
			if strings.Contains(e.DataType, "time.Time") {
				return "int64"
			} else {
				return e.DataType
			}
		}(),
		ColumnName:    e.Name.Source(),
		ColumnComment: e.Comment,
		Tag:           map[string]string{field.TagKeyJson: e.Name.Source()},
	}
}
