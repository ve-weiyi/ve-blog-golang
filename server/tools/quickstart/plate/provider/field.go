package provider

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/plate/field"
)

type Field struct {
	Name             string `json:"name"`
	FieldName        string `json:"field_name"`      // Field名
	FieldType        string `json:"field_type"`      // Field数据类型
	FieldJsonName    string `json:"field_jsonName"`  // FieldJson名
	FieldValueName   string `json:"field_valueName"` // Field值名
	ColumnComment    string `json:"column_comment"`  // 数据库字段描述
	MultilineComment bool
	Tag              field.Tag
	GORMTag          field.GormTag
	//FieldDefault    string `json:"fieldDefault"`    // 默认值
	//FieldDesc       string `json:"fieldDesc"`       // 中文名
	//FieldJson       string `json:"fieldJson"`       // FieldJson
	//DataType        string `json:"dataType"`        // 数据库字段类型(长度)
	//DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度

	//ColumnName      string `json:"columnName"`      // 数据库字段
	//FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	//DictType        string `json:"dictType"`        // 字典
	//Require         bool   `json:"require"`         // 是否必填
	//ErrorText       string `json:"errorText"`       // 校验失败文字
	//Clearable       bool   `json:"clearable"`       // 是否可清空
	//Sort            bool   `json:"sort"`            // 是否增加排序
}

// Tags ...
func (m *Field) Tags() string {
	if _, ok := m.Tag[field.TagKeyGorm]; ok {
		return m.Tag.Build()
	}

	if gormTag := strings.TrimSpace(m.GORMTag.Build()); gormTag != "" {
		m.Tag.Set(field.TagKeyGorm, gormTag)
	}
	return m.Tag.Build()
}

// GenType ...
func (m *Field) GenType() string {
	typ := strings.TrimLeft(m.FieldType, "*")
	switch typ {
	case "string", "bytes":
		return strings.Title(typ)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return strings.Title(typ)
	case "float64", "float32":
		return strings.Title(typ)
	case "bool":
		return strings.Title(typ)
	case "time.Time":
		return "Time"
	case "json.RawMessage", "[]byte":
		return "Bytes"
	default:
		return "Field"
	}
}

// EscapeKeyword escape keyword
func (m *Field) EscapeKeyword() *Field {
	return m.EscapeKeywordFor(GormKeywords)
}

// EscapeKeywordFor escape for specified keyword
func (m *Field) EscapeKeywordFor(keywords KeyWord) *Field {
	if keywords.FullMatch(m.FieldName) {
		m.FieldName += "_"
	}
	return m
}
