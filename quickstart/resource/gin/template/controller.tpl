package {{.Package}}

import (
{{- range .Imports}}
    {{.}}
{{- end}}
)

type {{.Name}}Controller struct {
	svcCtx *svc.ServiceContext
}

func New{{.Name}}Controller(svcCtx *svc.ServiceContext) *{{.Name}}Controller {
	return &{{.Name}}Controller{
		svcCtx: svcCtx,
	}
}

{{- range .Routes}}

// @Tags		{{$.Name}}
// @Summary		"{{.Doc}}"
// @Param		data	body		types.{{.Request}}		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.{{.Response}}}	"返回信息"
// @Router		{{.Path}} [{{.Method}}]
func (s *{{$.Name}}Controller) {{.Handler}}(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req types.{{.Request}}
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

    data, err := service.New{{$.Name}}Service(s.svcCtx).{{.Handler}}(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

{{- end}}
