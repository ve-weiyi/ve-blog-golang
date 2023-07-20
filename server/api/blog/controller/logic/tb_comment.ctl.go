package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// @Tags		Admin
// @Summary		获取用户评论列表
// @Security	ApiKeyUser
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo	true	"分页参数"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/comments [post]
func (s *AdminController) GetComments(c *gin.Context) {
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

	s.Log.JsonIndent(page)
	list, total, err := s.svcCtx.CommentService.FindCommonBackList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}
