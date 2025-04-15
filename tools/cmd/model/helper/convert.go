package helper

import (
	"fmt"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func ConvertTableToData(table *Table) ModelData {

	var fs []*ModelField
	for _, e := range table.Fields {
		fs = append(fs, &ModelField{
			Name:    jsonconv.Case2Camel(e.Name),
			Type:    strings.TrimPrefix(e.DataType, "u"),
			Tag:     fmt.Sprintf(`json:"%v" gorm:"column:%v"`, e.Name, e.Name),
			Comment: e.Comment,
		})
	}

	var ufs [][]*ModelField
	for _, es := range table.UniqueIndex {
		var u []*ModelField
		for _, e := range es {
			u = append(u, &ModelField{
				Name:    jsonconv.Case2Camel(e.Name),
				Type:    strings.TrimPrefix(e.DataType, "u"),
				Tag:     fmt.Sprintf(`json:"%v" gorm:"column:%v"`, e.Name, e.Name),
				Comment: e.Comment,
			})
		}
		ufs = append(ufs, u)
	}

	data := ModelData{
		TableName:           table.Name,
		UpperStartCamelName: jsonconv.Case2Camel(table.Name),
		LowerStartCamelName: jsonconv.FirstLower(jsonconv.Case2Camel(table.Name)),
		SnakeName:           jsonconv.Case2Snake(table.Name),
		Fields:              fs,
		UniqueFields:        ufs,
	}

	return data
}
