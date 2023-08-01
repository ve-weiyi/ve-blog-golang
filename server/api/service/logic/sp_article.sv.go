package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	request2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 根据id获取Article记录
func (s *ArticleService) GetArticleDetails(reqCtx *request2.Context, id int) (data *response.ArticleDetails, err error) {
	// 查询id对应文章
	article, err := s.svcCtx.ArticleRepository.FindArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, err := s.svcCtx.CategoryRepository.FindCategory(reqCtx, article.CategoryID)
	if err != nil {
		return nil, err
	}

	// 查询文章标签
	tags, err := s.svcCtx.TagRepository.FindArticleTagList(article.ID)
	if err != nil {
		return nil, err
	}

	// 查询推荐文章
	rmArticle, err := s.svcCtx.ArticleRepository.FindRecommendArticle(article.CategoryID)
	if err != nil {
		return nil, err
	}
	// 查询最新文章
	page := &request2.PageQuery{
		Page:     0,
		PageSize: 5,
		Sorts: []*request2.Sort{
			{Field: "id", Order: "desc"},
		},
	}
	newestArticle, _, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, err
	}
	// 查询上一篇文章
	lastArticle, err := s.svcCtx.ArticleRepository.FindLastArticle(id)
	if err != nil {
		return nil, err
	}
	// 查询下一篇文章
	nextArticle, err := s.svcCtx.ArticleRepository.FindNextArticle(id)
	if err != nil {
		return nil, err
	}

	resp := convertResponseArticle(article)
	resp.CategoryName = category.CategoryName
	resp.ArticleTagList = convertTagList(tags)
	resp.RecommendArticleList = convertRecommendArticles(rmArticle)
	resp.NewestArticleList = convertRecommendArticles(newestArticle)
	resp.LastArticle = convertArticlePagination(lastArticle)
	resp.NextArticle = convertArticlePagination(nextArticle)
	return resp, nil
}

// 分页获取Article记录
func (s *ArticleService) GetArticleList(reqCtx *request2.Context, page *request2.PageQuery) (list []*response.ArticleDTO, total int64, err error) {
	// 查询文章列表
	articles, total, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.FindCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.FindArticleTagList(article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}
	return list, total, err
}

func (s *ArticleService) GetArticleListByCondition(reqCtx *request2.Context, req *request2.ArticleCondition) (data *response.ArticleConditionDTO, err error) {
	resp := &response.ArticleConditionDTO{}

	// 查询文章列表
	var articles []*entity.Article

	if req.CategoryID != 0 {
		category, err := s.svcCtx.CategoryRepository.FindCategory(reqCtx, req.CategoryID)
		if err != nil {
			return nil, err
		}
		articles, _, err = s.svcCtx.ArticleRepository.FindArticleListByCategoryId(category.ID)
		resp.ConditionName = category.CategoryName
	} else if req.TagID != 0 {
		tag, err := s.svcCtx.TagRepository.FindTag(reqCtx, req.TagID)
		if err != nil {
			return nil, err
		}
		articles, _, err = s.svcCtx.ArticleRepository.FindArticleListByTagId(tag.ID)
		resp.ConditionName = tag.TagName
	}

	var list []*response.ArticleDTO
	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.FindCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.FindArticleTagList(article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}

	resp.ArticleDTOList = list
	return resp, err
}

func (s *ArticleService) GetArticleArchives(reqCtx *request2.Context, page *request2.PageQuery) (list []*response.ArticleRecommendDTO, total int64, err error) {
	page.Sorts = []*request2.Sort{
		{Field: "id", Order: "desc"},
	}
	newestArticle, total, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	return convertRecommendArticles(newestArticle), total, err
}
