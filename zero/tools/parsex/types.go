package parsex

type (
	// Table describes a mysql table
	Table struct {
		Name        string
		Db          string
		PrimaryKey  Primary
		UniqueIndex map[string][]*Field
		Fields      []*Field
	}

	// Primary describes a primary key
	Primary struct {
		Field
		AutoIncrement bool
	}

	// Field describes a table field
	Field struct {
		//NameOriginal    string
		Name string
		//ThirdPkg        string
		DataType        string
		Comment         string
		SeqInIndex      int
		OrdinalPosition int
		//ContainsPQ      bool
	}

	// KeyType types alias of int
	KeyType int
)

type Route struct {
	Doc      string // 用户接口
	Handler  string // UserHandler
	Method   string // POST、GET、PUT、DELETE
	Path     string // /api/v1/user
	Request  string
	Response string
}

type GroupRoute struct {
	Name   string // account
	Routes []Route
}

//// 文件结构
//type FileSpec struct {
//	Package string
//	Imports []string
//	Structs []StructSpec
//}
//
//// 结构体
//type StructSpec struct {
//	Name      string
//	Fields    []FieldSpec
//	Functions []FunctionSpec
//}
//
//// 字段
//type FieldSpec struct {
//	Name string
//	Type string
//}
//
//// 函数
//type FunctionSpec struct {
//	FuncName  string
//	ParamsIn  []FieldSpec
//	ParamsOut []FieldSpec
//	Data      any
//}
