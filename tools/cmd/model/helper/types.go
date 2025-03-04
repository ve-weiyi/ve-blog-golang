package helper

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
		Name    string // 属性名称  Name
		Type    string // 属性类型  string、int、bool、float、{UpperStartCamelName}
		Tag     string // json tag
		Comment string // 行尾注释
	}
)
