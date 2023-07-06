package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type ArticleController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewArticleController(svcCtx *svc.ControllerContext) *ArticleController {
	return &ArticleController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Article
// @Summary	创建文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Article							true	"创建文章"
// @Success	200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/create [post]
func (s *ArticleController) CreateArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBindJSON(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.CreateArticle(reqCtx, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	删除文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Article		true	"删除文章"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/article/delete [delete]
func (s *ArticleController) DeleteArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBindJSON(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.DeleteArticle(reqCtx, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	更新文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Article							true	"更新文章"
// @Success	200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/update [put]
func (s *ArticleController) UpdateArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBindJSON(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.UpdateArticle(reqCtx, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	用id查询文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Article							true	"用id查询文章"
// @Success	200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/find [post]
func (s *ArticleController) FindArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBindJSON(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.GetArticleDetails(reqCtx, article.ID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	批量删除文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		[]int				true	"批量删除文章"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/article/deleteByIds [delete]
func (s *ArticleController) DeleteArticleByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBindJSON(c, &IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.DeleteArticleByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	分页获取文章列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		page	body		request.PageInfo													true	"分页获取文章列表"
// @Success	200		{object}	response.Response{data=response.PageResult{list=[]entity.Article}}	"返回信息"
// @Router		/article/list [post]
func (s *ArticleController) GetArticleList(c *gin.Context) {
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

	list, total, err := s.svcCtx.ArticleService.GetArticleList(reqCtx, &page)
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

// @Tags		Article
// @Summary	分页获取文章列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		page	body		request.PageInfo													true	"分页获取文章列表"
// @Success	200		{object}	response.Response{data=response.PageResult{list=[]entity.Article}}	"返回信息"
// @Router		/article/condition [post]
func (s *ArticleController) GetArticleListByCondition(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.ArticleCondition
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.GetArticleListByCondition(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary	更新文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Article							true	"更新文章"
// @Success	200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/like [put]
func (s *ArticleController) LikeArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBindJSON(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.UpdateArticle(reqCtx, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
