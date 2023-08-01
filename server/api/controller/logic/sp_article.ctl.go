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
func (s *ArticleController) GetArticleDetails(c *gin.Context) {
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

	data, err := s.svcCtx.ArticleService.GetArticleDetails(reqCtx, id)
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

	data, err := s.svcCtx.ArticleService.GetArticleDetails(reqCtx, id)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Article
// @Summary		文章归档
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/archives [get]
func (s *ArticleController) GetArticleArchives(c *gin.Context) {
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

	list, total, err := s.svcCtx.ArticleService.GetArticleArchives(reqCtx, &page)
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
// @Summary		分页获取文章列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.PageQuery													true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Article}}	"返回信息"
// @Router		/article/list/condition [post]
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
