package controller

import (
	"github.com/gin-gonic/gin"

)

type {{Case2Camel .Name}}Controller struct {
	controller.BaseController
	svcCtx *svc.ServiceContext
}

func New{{Case2Camel .Name}}Controller(svcCtx *svc.ServiceContext) *{{Case2Camel .Name}}Controller {
	return &{{Case2Camel .Name}}Controller{
		svcCtx: svcCtx,
		BaseController: controller.NewBaseController(),
	}
}

{{- range .Routes}}

// @Tags		{{$.Name}}
// @Summary		{{.Doc}}
// @Param		data	body		types.{{.Request}}		true	"请求参数"
// @Success		200		{object}	response.Response{data=types.{{.Response}}}	"返回信息"
// @Router		{{.Path}} [{{.Method}}]
func (s *{{Case2Camel $.Name}}Controller) {{.Handler}}(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req types.{{.Request}}
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

    data, err := service.New{{Case2Camel $.Name}}Service(s.svcCtx).{{.Handler}}(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

{{- end}}
