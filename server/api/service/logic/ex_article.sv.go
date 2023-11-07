package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 根据id获取Article记录
func (s *ArticleService) FindArticleDetails(reqCtx *request.Context, id int) (data *response.ArticlePaginationDTO, err error) {
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
	tags, err := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)
	if err != nil {
		return nil, err
	}

	// 查询推荐文章
	rmArticle, err := s.svcCtx.ArticleRepository.FindRecommendArticle(reqCtx, article.CategoryID)
	if err != nil {
		return nil, err
	}
	// 查询最新文章
	page := &request.PageQuery{
		Page:     0,
		PageSize: 5,
		Sorts: []*request.Sort{
			{Field: "id", Order: "desc"},
		},
	}
	newestArticle, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, err
	}
	// 查询上一篇文章
	lastArticle, err := s.svcCtx.ArticleRepository.FindLastArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}
	// 查询下一篇文章
	nextArticle, err := s.svcCtx.ArticleRepository.FindNextArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}

	resp := convertResponseArticle(article)
	resp.CategoryName = category.CategoryName
	resp.ArticleTagList = convertTagList(tags)
	resp.RecommendArticleList = convertArticlePreviewList(rmArticle)
	resp.NewestArticleList = convertArticlePreviewList(newestArticle)
	resp.LastArticle = convertArticlePreview(lastArticle)
	resp.NextArticle = convertArticlePreview(nextArticle)
	return resp, nil
}

// 分页获取Article记录
func (s *ArticleService) FindArticleDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticleDetails, total int64, err error) {
	// 查询文章列表
	articles, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.FindCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}
	return list, total, err
}

func (s *ArticleService) FindArticleListByCondition(reqCtx *request.Context, req *request.ArticleCondition) (data *response.ArticleConditionDTO, err error) {
	data = &response.ArticleConditionDTO{}
	// 查询文章列表
	var articles []*entity.Article

	if req.CategoryID != 0 {
		category, err := s.svcCtx.CategoryRepository.FindCategory(reqCtx, req.CategoryID)
		if err != nil {
			return nil, err
		}
		articles, err = s.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, category.ID)
		data.ConditionName = category.CategoryName
	} else if req.TagID != 0 {
		tag, err := s.svcCtx.TagRepository.FindTag(reqCtx, req.TagID)
		if err != nil {
			return nil, err
		}
		articles, err = s.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, tag.ID)
		data.ConditionName = tag.TagName
	}

	var list []*response.ArticleDetails
	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.FindCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}

	data.ArticleDTOList = list
	return data, err
}

func (s *ArticleService) FindArticleArchives(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticlePreviewDTO, total int64, err error) {
	// 查找最新数据
	page.Sorts = []*request.Sort{
		{Field: "id", Order: "desc"},
	}
	newestArticle, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return convertArticlePreviewList(newestArticle), total, err
}
