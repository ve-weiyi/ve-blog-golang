package parser

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/tools/meta"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func TestParsePlainText(t *testing.T) {
	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := os.WriteFile(sqlFile, []byte("plain text"), 0o777)
	assert.Nil(t, err)

	_, err = parser.Parse(sqlFile, "go_zero", false)
	assert.NotNil(t, err)
}

func TestParseSelect(t *testing.T) {
	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := os.WriteFile(sqlFile, []byte("select * from user"), 0o777)
	assert.Nil(t, err)

	tables, err := parser.Parse(sqlFile, "go_zero", false)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(tables))
}

//go:embed testdata/user.sql
var user string

func TestConvertColumn(t *testing.T) {
	t.Run("missingPrimaryKey", func(t *testing.T) {
		columnData := model.ColumnData{
			Db:    "user",
			Table: "user",
			Columns: []*model.Column{
				{
					DbColumn: &model.DbColumn{
						Name:     "id",
						DataType: "bigint",
					},
				},
			},
		}
		_, err := columnData.Convert()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "missing primary key")
	})

	t.Run("jointPrimaryKey", func(t *testing.T) {
		columnData := model.ColumnData{
			Db:    "user",
			Table: "user",
			Columns: []*model.Column{
				{
					DbColumn: &model.DbColumn{
						Name:     "id",
						DataType: "bigint",
					},
					Index: &model.DbIndex{
						IndexName: "PRIMARY",
					},
				},
				{
					DbColumn: &model.DbColumn{
						Name:     "mobile",
						DataType: "varchar",
						Comment:  "手机号",
					},
					Index: &model.DbIndex{
						IndexName: "PRIMARY",
					},
				},
			},
		}
		_, err := columnData.Convert()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "joint primary key is not supported")
	})

	t.Run("normal", func(t *testing.T) {
		columnData := model.ColumnData{
			Db:    "user",
			Table: "user",
			Columns: []*model.Column{
				{
					DbColumn: &model.DbColumn{
						Name:     "id",
						DataType: "bigint",
						Extra:    "auto_increment",
					},
					Index: &model.DbIndex{
						IndexName:  "PRIMARY",
						SeqInIndex: 1,
					},
				},
				{
					DbColumn: &model.DbColumn{
						Name:     "mobile",
						DataType: "varchar",
						Comment:  "手机号",
					},
					Index: &model.DbIndex{
						IndexName:  "mobile_unique",
						SeqInIndex: 1,
					},
				},
			},
		}

		table, err := columnData.Convert()
		assert.Nil(t, err)
		assert.True(t, table.PrimaryKey.Index.IndexName == "PRIMARY" && table.PrimaryKey.Name == "id")
		for _, item := range table.Columns {
			if item.Name == "mobile" {
				assert.True(t, item.Index.NonUnique == 0)
				break
			}
		}
	})
}

//go:embed tpl/api.tpl
var apiTpl string

//go:embed tpl/rpc.tpl
var rpcTpl string

//go:embed tpl/model.tpl
var modelTpl string

func TestParseSql(t *testing.T) {
	tables, err := parser.Parse("testdata/test.sql", "go_zero", false)
	assert.Nil(t, err)

	//log.Printf("table %+v", jsonconv.ObjectToJsonIndent(tables))

	for _, table := range tables {
		log.Printf("%+v", table.Name)

		metas := NewMetas(table)

		for _, meta := range metas {
			err = meta.Execute()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func NewMetas(table *parser.Table) []invent.TemplateMeta {
	var fs []*meta.Field
	for _, e := range table.Fields {
		//log.Printf("%+v", jsonconv.ObjectToJsonIndent(e))

		fs = append(fs, convertField(e))
	}

	data := map[string]any{
		"DbName":              table.Db.Source(),
		"TableName":           table.Name.Source(),
		"UpperStartCamelName": jsonconv.Case2Camel(table.Name.Source()),
		"LowerStartCamelName": jsonconv.Case2CamelLowerStart(table.Name.Source()),
		"SnakeName":           jsonconv.Camel2Case(table.Name.Source()),
		"CommentName":         table.Name.Source(),
		"Fields":              fs,
		"ImportPkgPaths":      nil,
	}

	var metas []invent.TemplateMeta
	metas = append(metas, invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fmt.Sprintf("./test/api/%v.api", table.Name.Source()),
		TemplateString: apiTpl,
		FunMap:         nil,
		Data:           data,
	})

	metas = append(metas, invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fmt.Sprintf("./test/rpc/%v.proto", table.Name.Source()),
		TemplateString: rpcTpl,
		FunMap: map[string]any{
			"add": func(index int, num int) int {
				return index + num
			},
		},
		Data: data,
	})

	metas = append(metas, invent.TemplateMeta{
		Key:            "",
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fmt.Sprintf("./test/model/%v_model.go", table.Name.Source()),
		TemplateString: modelTpl,
		FunMap:         nil,
		Data:           data,
	})

	return metas
}

func convertField(e *parser.Field) *meta.Field {

	return &meta.Field{
		Name:    jsonconv.Case2Camel(e.Name.Source()),
		Type:    e.DataType,
		Comment: e.Comment,
		Tag: map[string][]string{
			"json": []string{
				e.Name.Source(),
			},
		},
		Docs:     nil,
		IsInline: false,
	}
}
