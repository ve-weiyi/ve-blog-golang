package mysql

import (
	"fmt"
	"os"
	"path"

	"github.com/ve-weiyi/pkg/kit/quickstart/gotplgen"
	"github.com/ve-weiyi/pkg/utils/filex"
	"github.com/ve-weiyi/pkg/utils/jsonconv"
)

type (
	ModelData struct {
		TableName           string
		UpperStartCamelName string
		LowerStartCamelName string
		SnakeName           string
		Fields              []*ModelField
		UniqueFields        [][]*ModelField
	}

	ModelField struct {
		Name string // 属性名称  Name
		Type string // 属性类型  string、int、bool、float、{UpperStartCamelName}

		Tag     string // json tag
		Comment string // 行尾注释
	}
)

func generateModel(models []*ModelData, tplPath string, outPath string, nameAs string) error {
	var metas []gotplgen.TemplateMeta

	tpl, err := os.ReadFile(filex.ToAbs(tplPath))
	if err != nil {
		return err
	}

	for _, model := range models {
		meta := gotplgen.TemplateMeta{
			Mode:           gotplgen.ModeCreateOrReplace,
			CodeOutPath:    path.Join(outPath, fmt.Sprintf(nameAs, model.TableName)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"funcFieldsKey": func(fs []*ModelField) string {
					var name string
					for _, ff := range fs {
						name += ff.Name
					}
					return name
				},
				"funcFieldsKeyVar": func(fs []*ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						tp := ff.Type
						if name != "" {
							name += ", "
						}
						name += fmt.Sprintf("%s %s", v, extractBaseType(tp))
					}
					return name
				},
				"funcFieldsKeyCond": func(fs []*ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += " and "
						}
						name += fmt.Sprintf("`%s` = ?", v)
					}
					return name
				},
				"funcFieldsKeyCondVar": func(fs []*ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += ", "
						}
						name += v
					}
					return name
				},
			},
			Data: model,
		}
		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

// extractBaseType 提取基础类型原型
// sql.NullString -> string
// sql.NullInt64 -> int64
// sql.NullTime -> time.Time
// string -> string (保持不变)
func extractBaseType(dataType string) string {
	switch dataType {
	case "sql.NullString":
		return "string"
	case "sql.NullInt64":
		return "int64"
	case "sql.NullInt32":
		return "int32"
	case "sql.NullFloat64":
		return "float64"
	case "sql.NullBool":
		return "bool"
	case "sql.NullTime":
		return "time.Time"
	default:
		return dataType
	}
}
