package mysql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/pkg/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/dbparser"
)

func ConvertTableToData(table *dbparser.Table) *ModelData {

	var fs []*ModelField
	for _, e := range table.Fields {
		fs = append(fs, ConvertField(e))
	}

	var ufs [][]*ModelField
	// 先收集所有 key
	keys := make([]string, 0, len(table.UniqueIndex))
	for k := range table.UniqueIndex {
		keys = append(keys, k)
	}

	// 排序 key
	sort.Strings(keys)

	// 遍历排序后的 key，直接构建 ufs
	for _, k := range keys {
		es := table.UniqueIndex[k]
		u := make([]*ModelField, 0, len(es))
		for _, e := range es {
			u = append(u, ConvertField(e))
		}
		ufs = append(ufs, u)
	}

	data := &ModelData{
		TableName:           table.Name,
		UpperStartCamelName: jsonconv.Case2Camel(table.Name),
		LowerStartCamelName: jsonconv.FirstLower(jsonconv.Case2Camel(table.Name)),
		SnakeName:           jsonconv.Case2Snake(table.Name),
		Fields:              fs,
		UniqueFields:        ufs,
	}

	return data
}

func ConvertField(f *dbparser.Field) *ModelField {
	return &ModelField{
		Name:    jsonconv.Case2Camel(f.Name),
		Type:    strings.TrimPrefix(f.DataType, "u"),
		Tag:     fmt.Sprintf(`json:"%v" gorm:"column:%v"`, f.Name, f.Name),
		Comment: f.Comment,
	}
}
