package apiparser

// api方法定义信息
type ApiDeclare struct {
	Tag          string `json:"tag"`
	FunctionName string `json:"function_name"`
	Summary      string `json:"summary"`

	//Base     string
	Router   string      `json:"router"`
	Method   string      `json:"method,omitempty"`
	Header   []*ApiParam `json:"header,omitempty"` // 请求头参数
	Path     []*ApiParam `json:"path,omitempty"`   // 路径参数 path
	Query    []*ApiParam `json:"query,omitempty"`  // 查询参数 query
	Form     []*ApiParam `json:"form,omitempty"`   // 表单参数 form-data
	Body     *ApiParam   `json:"body,omitempty"`   // 请求体参数
	Response string      `json:"response"`         // 响应参数
}

// 参数定义
type ApiParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ApiParser interface {
	ParseApiDocsByRoot(root ...string) ([]*ApiDeclare, error)
}
