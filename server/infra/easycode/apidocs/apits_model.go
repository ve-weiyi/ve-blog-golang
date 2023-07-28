package apidocs

type ApiDoc struct {
	Tag      string
	Function map[string]*ApiDeclare
}

type ApiCommentLine struct {
	Tag     string
	Content string
}

// api方法定义信息
type ApiDeclare struct {
	Tag          string
	FunctionName string
	Summary      string

	//Base     string
	Url      string
	Method   string
	Header   []*ApiParam // 请求头参数
	Path     []*ApiParam // 路径参数 path
	Query    []*ApiParam // 查询参数 query
	Form     []*ApiParam // 表单参数 form-data
	Body     *ApiParam   // 请求体参数
	Response string      // 响应参数
}

type ApiParam struct {
	Name string
	Type string
}

// model定义信息
type ModelDeclare struct {
	Pkg    string
	Name   string
	Extend *ModelDeclare
	Fields []*ModelField
}

type ModelField struct {
	Name    string // 属性名称
	Type    string // 属性类型
	Comment string // 属性的注释
}

// api文档信息
type TsApiDoc struct {
	Tag            string            // 标签、分类
	ImportPkgPaths []string          // 导入的包
	ModelDeclares  []*TsModelDeclare // 导入的model
	ApiDeclares    []*TsApiDeclare   // 所有api方法
}

type TsApiDeclare struct {
	Tag          string
	FunctionName string
	Summary      string

	//Base     string
	Url      string
	Method   string
	Header   []*ApiParam // 请求头参数
	Path     []*ApiParam // 路径参数 path
	Query    []*ApiParam // 查询参数 query
	Form     []*ApiParam // 表单参数 form-data
	Body     *ApiParam   // 请求体参数
	Response string      // 响应参数
}

type TsModelDeclare struct {
	Name   string
	Extend *TsModelDeclare
	Fields []*ModelField
}
