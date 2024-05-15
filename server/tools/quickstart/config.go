package quickstart

import (
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/gorm_parser/model"
)

type Config struct {
	DbEngin *gorm.DB

	model.FieldConfig

	ReplaceMode int //是否替换文件 0:创建或替换 1:只创建 2:只替换

	OutPath   string                                   // 输出路径
	OutFileNS func(tableName string) (fileName string) // 输出文件名称

	//FieldNameNS  func(columnName string) (fieldName string) // 转换数据库字段名为go字段名
	//FieldJsonNS  func(columnName string) (jsonName string) // 转换数据库字段名为go json tag字段名
	//FieldValueNS func(columnName string) (valueName string) // 转换数据库字段名为go字段名

	IsIgnoreKey func(key string) bool // 是否忽略key
}
