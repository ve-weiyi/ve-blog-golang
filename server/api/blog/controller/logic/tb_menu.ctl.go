package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// @Tags		Menu
// @Summary		获取菜单列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Menu}	"返回信息"
// @Router		/admin/menus [post]
func (s *MenuController) GetMenus(c *gin.Context) {
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

	list, total, err := s.svcCtx.MenuService.GetAllMenusList(reqCtx, &page)
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
