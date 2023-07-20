package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type AdminController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewAdminController(svcCtx *svc.ControllerContext) *AdminController {
	return &AdminController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

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

// @Tags		Admin
// @Summary		获取角色列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Role}	"返回信息"
// @Router		/admin/roles [post]
func (s *AdminController) GetRoles(c *gin.Context) {
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

	list, total, err := s.svcCtx.RoleService.GetRoles(reqCtx, &page)
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

// @Tags		Admin
// @Summary		获取菜单列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Menu}	"返回信息"
// @Router		/admin/menus [post]
func (s *AdminController) GetMenus(c *gin.Context) {
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

// @Tags		Admin
// @Summary		获取api列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo						true	"分页参数"
// @Success		200		{object}	response.Response{data=[]entity.Api}	"返回信息"
// @Router		/admin/apis [post]
func (s *AdminController) GetApis(c *gin.Context) {
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
