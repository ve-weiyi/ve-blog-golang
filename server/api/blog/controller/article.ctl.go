package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type ArticleController struct {
	svcCtx *svc.ServiceContext
}

func NewArticleController(svcCtx *svc.ServiceContext) *ArticleController {
	return &ArticleController{
		svcCtx: svcCtx,
	}
}

// @Tags		Article
// @Summary		保存文章
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.ArticleDetailsDTOReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/admin/save_article [post]
func (s *ArticleController) SaveArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var article dto.ArticleDetailsDTOReq
	err = request.ShouldBind(c, &article)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).SaveArticle(reqCtx, &article)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		删除文章
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	request		body		request.IdReq							true	"Article.id"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}			"返回信息"
// @Router		/admin/delete_article [delete]
func (s *ArticleController) DeleteArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).DeleteArticle(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Article
// @Summary		查询文章
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	request		body		request.IdReq							true	"Article.id"
// @Success		200		{object}	response.Body{data=dto.ArticleBack}	"返回信息"
// @Router 		/admin/find_article [post]
func (s *ArticleController) FindArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).FindArticle(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Article
// @Summary		分页获取文章列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.PageQuery 			true 	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.ArticleBack}}	"返回信息"
// @Router		/admin/article/find_article_list [post]
func (s *ArticleController) FindArticleList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewArticleService(s.svcCtx).FindArticleList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags 	 	Article
// @Summary		删除文章-逻辑删除
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.ArticleDeleteReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router 		/admin/article/delete [put]
func (s *ArticleController) LogicDeleteArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ArticleDeleteReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).UpdateArticleDelete(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Article
// @Summary		更新文章
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		request.ArticleTopReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router 		/admin/article/top [put]
func (s *ArticleController) TopArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ArticleTopReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).UpdateArticleTop(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		文章归档(时间轴)
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页获取文章列表"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.ArticlePreviewDTO}}	"返回信息"
// @Router		/article/archives [post]
func (s *ArticleController) FindArticleArchives(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewArticleService(s.svcCtx).FindArticleArchives(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Article
// @Summary		通过标签或者id获取文章列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.ArticleClassifyReq											true	"分页获取文章列表"
// @Success		200		{object}	response.Body{data=dto.ArticleClassifyResp}	"返回信息"
// @Router		/article/article_classify_category [post]
func (s *ArticleController) FindArticleClassifyCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ArticleClassifyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).FindArticleClassifyCategory(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		通过标签或者id获取文章列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.ArticleClassifyReq											true	"分页获取文章列表"
// @Success		200		{object}	response.Body{data=dto.ArticleClassifyResp}	"返回信息"
// @Router		/article/article_classify_tag [post]
func (s *ArticleController) FindArticleClassifyTag(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req dto.ArticleClassifyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).FindArticleClassifyTag(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		文章相关推荐
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param 	 	request		body		request.IdReq										true	"Article.id"
// @Success		200		{object}	response.Body{data=dto.ArticlePageDetailsDTO}	"返回信息"
// @Router		/article/find_article_recommend [post]
func (s *ArticleController) FindArticleRecommend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).FindArticleRecommend(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Article
// @Summary		分页获取文章列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.PageQuery 			true 	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.ArticleHome}}	"返回信息"
// @Router		/article/find_article_list [post]
func (s *ArticleController) FindArticleHomeList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewArticleService(s.svcCtx).FindArticleHomeList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Article
// @Summary		点赞文章
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		req		body		request.IdReq										true	"Article.id"
// @Success		200		{object}	response.Body{data=entity.Article}	"返回信息"
// @Router		/article/like_article [put]
func (s *ArticleController) LikeArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewArticleService(s.svcCtx).LikeArticle(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
