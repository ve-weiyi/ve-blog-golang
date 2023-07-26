package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// @Tags		Article
// @Summary		文章归档
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo													true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
// @Router		/article/archives [get]
func (s *ArticleController) GetArticleArchives(c *gin.Context) {
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
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageInfo													true	"分页获取文章列表"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Article}}	"返回信息"
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
// @Summary		更新文章
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Article							true	"更新文章"
// @Success		200		{object}	response.Response{data=entity.Article}	"返回信息"
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
