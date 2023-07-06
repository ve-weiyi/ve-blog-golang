package apidocs

type ApiTs struct {
	FileName string
	Function map[string]*ApiTsMethod
}

type ApiTsMethod struct {
	Description string
	Method      string
	Url         string
	Header      []string // 请求头参数
	Path        []string // 路径参数
	Params      []string // 请求query和form-data参数
	Body        string   // 请求体参数
}
