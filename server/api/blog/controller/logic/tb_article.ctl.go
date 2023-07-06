package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
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
// @Summary		创建文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Article							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/create [post]
func (s *ArticleController) CreateArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBind(c, &article)
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
// @Summary		删除文章
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.Article 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/article/delete [delete]
func (s *ArticleController) DeleteArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBind(c, &article)
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

// @Tags 	 	Article
// @Summary		更新文章
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.Article							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router 		/article/update [put]
func (s *ArticleController) UpdateArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBind(c, &article)
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

// @Tags 	 	Article
// @Summary		查询文章
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		body		entity.Article							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.Article}	"返回信息"
// @Router 		/article/query [get]
func (s *ArticleController) GetArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = s.ShouldBind(c, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.GetArticle(reqCtx, &article)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Article
// @Summary		批量删除文章
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/article/deleteByIds [delete]
func (s *ArticleController) DeleteArticleByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBind(c, &IDS)
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

// @Tags 	 	Article
// @Summary		分页获取文章列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Article}}	"返回信息"
// @Router		/article/list [post]
func (s *ArticleController) FindArticleList(c *gin.Context) {
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

	list, total, err := s.svcCtx.ArticleService.FindArticleList(reqCtx, &page)
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
