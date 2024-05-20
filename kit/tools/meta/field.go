package meta

import (
	"strings"
)

// Field describes the field of a structure
type Field struct {
	Name     string
	Type     string              // 数据类型字面值，如：string、map[int]string、[]int64、[]*User
	Tag      map[string][]string // 标签，如：json、gorm、protobuf  `json:"tag,omitempty"`
	Comment  string
	Docs     []string // 成员头顶注释说明
	IsInline bool     //是否是内联结构体
}

func (m *Field) Tags() string {
	var tags string

	// json  [tag,omitempty]
	for k, v := range m.Tag {
		tags += k
		if len(v) > 0 {
			tags += ":" + strings.Join(v, ",")
		}
		tags += " "
	}

	return tags
}
