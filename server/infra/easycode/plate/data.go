package plate

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate/provider"
)

//type CodeStructMeta interface {
//	Package()   string
//	TableName()string
//	StructName()string
//	ValueName()string
//	JsonName() string
//	StructComment() string
//	ImportPkgPaths() []string
//}

// AutoCodeStructData 初始版本自动化代码工具
type AutoCodeStructData struct {
	Package        string            `json:"package"`
	TableName      string            `json:"tableName"`   // 表名 				auto_code
	StructName     string            `json:"structName"`  // Struct名称 		AutoCode 大写驼峰命名
	ValueName      string            `json:"valueName"`   // Struct变量名 		autoCode 小写驼峰命名
	JsonName       string            `json:"jsonName"`    // StructJson名		auto_code api路径前缀
	StructComment  string            `json:"chineseName"` // Struct中文名称 	创建api的描述和注释
	Fields         []*provider.Field `json:"fields,omitempty"`
	ImportPkgPaths []string
}

//func (m *AutoCodeStructData) Reverse() *AutoCodeStructData {
//	if m.TableName == "" {
//		return nil
//	}
//
//	tableName := m.TableName
//	if m.Package == "" {
//		m.Package = jsonconv.Case2CamelNotFirst(tableName)
//	}
//
//	if m.StructName == "" {
//		m.StructName = jsonconv.Case2Camel(tableName)
//	}
//
//	if m.ValueName == "" {
//		m.ValueName = jsonconv.Case2CamelNotFirst(tableName)
//	}
//
//	if m.JsonName == "" {
//		m.Package = jsonconv.Camel2Case(tableName)
//	}
//
//	if m.Fields == nil {
//		m.Fields = []*provider.Field{
//			&provider.Field{
//				FieldName: "ID",
//				FieldType: "int",
//				Tag: map[string]string{
//					"json": "id",
//				},
//				GORMTag: field.GormTag{
//					"column":     []string{"id"},
//					"type":       []string{"int"},
//					"primaryKey": []string{"autoIncrement:true"},
//				},
//				ColumnComment:    "ID",
//				MultilineComment: false,
//			},
//		}
//	}
//	return m
//}
