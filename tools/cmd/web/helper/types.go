package helper

type TsApiGroup struct {
	Name       string // account
	Prefix     string
	Middleware []string
	Routes     []TsApiRoute
}

type TsApiRoute struct {
	Summery  string
	Handler  string
	Path     string
	Method   string
	Request  string
	Response string
}

type TsType struct {
	Comment string
	Name    string
	Extends []string
	Fields  []TsTypeField
}

type TsTypeField struct {
	Name     string
	Type     string
	Comment  string
	Nullable bool
}
