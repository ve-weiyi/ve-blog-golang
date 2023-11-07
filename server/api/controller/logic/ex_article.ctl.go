package logic

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// @Tags		Article
// @Summary		文章详情
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		id		path		int										true	"Article id"
// @Success		200		{object}	response.Response{data=response.ArticleDetails}	"返回信息"
// @Router		/article/{id}/details [get]
func (s *ArticleController) FindArticleDetails(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.FindArticleDetails(reqCtx, id)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		分页获取文章详情列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页信息"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.ArticlePaginationDTO}}	"返回信息"
// @Router		/article/list/details [post]
func (s *ArticleController) FindArticleDetailsList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ArticleService.FindArticleDetailsList(reqCtx, &page)
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
// @Summary		文章归档(时间轴)
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.ArticlePreviewDTO}}	"返回信息"
// @Router		/article/archives [post]
func (s *ArticleController) FindArticleArchives(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ArticleService.FindArticleArchives(reqCtx, &page)
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
// @Summary		通过标签或者id获取文章列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.ArticleCondition											true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=response.ArticleConditionDTO}	"返回信息"
// @Router		/article/list/condition [post]
func (s *ArticleController) FindArticleListByCondition(c *gin.Context) {
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

	data, err := s.svcCtx.ArticleService.FindArticleListByCondition(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		点赞文章
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		id		path		int										true	"Article id"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/{id}/like [put]
func (s *ArticleController) LikeArticle(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ArticleService.FindArticleDetails(reqCtx, id)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
