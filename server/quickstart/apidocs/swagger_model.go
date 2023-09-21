package apidocs

// SwaggerDefinition 包含 Swagger 文档的定义
type SwaggerDefinition struct {
	Swagger  string           `json:"swagger"`
	Info     SwaggerInfo      `json:"info"`
	Host     string           `json:"host"`
	BasePath string           `json:"basePath"`
	Paths    map[string]Paths `json:"paths"`
}

// SwaggerInfo 包含 Swagger 文档的基本信息
type SwaggerInfo struct {
	Description       string         `json:"description"`
	Title             string         `json:"title"`
	TermsOfServiceURL string         `json:"termsOfServiceUrl"`
	Contact           SwaggerContact `json:"contact"`
	License           SwaggerLicense `json:"license"`
	Version           string         `json:"version"`
}

// SwaggerContact 包含 Swagger 文档的联系人信息
type SwaggerContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

// SwaggerLicense 包含 Swagger 文档的许可信息
type SwaggerLicense struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Paths 包含 Swagger 文档的路径定义
type Paths map[string]PathItem

// PathItem 包含 Swagger 文档的路径项定义
type PathItem struct {
	Tags        []string            `json:"tags"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	Security    []interface{}       `json:"security"`
	Consumes    []string            `json:"consumes"`
	Produces    []string            `json:"produces"`
	Parameters  []Parameter         `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
}

// Parameter 包含 Swagger 文档的参数定义
type Parameter struct {
	Name        string                 `json:"name"`
	In          string                 `json:"in"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Schema      map[string]interface{} `json:"schema"`
}

// Response 包含 Swagger 文档的响应定义
type Response struct {
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
}
