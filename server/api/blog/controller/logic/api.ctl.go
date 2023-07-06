package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
)

// @Tags		Api
// @Summary	获取api列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Api							true	"创建api路由"
// @Success	200		{object}	response.Response{data=entity.Api}	"返回信息"
// @Router		/admin/apis [post]
func (s *ApiController) GetApiTreeList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ApiService.GetAllApiList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}
