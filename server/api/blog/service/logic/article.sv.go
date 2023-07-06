package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// 根据id获取Article记录
func (s *ArticleService) GetArticleDetails(reqCtx *request.Context, id int) (data *response.ArticleDetails, err error) {
	// 查询id对应文章
	article, err := s.svcCtx.ArticleRepository.GetArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, err := s.svcCtx.CategoryRepository.GetCategory(reqCtx, article.CategoryID)
	if err != nil {
		return nil, err
	}

	// 查询文章标签
	tags, err := s.svcCtx.TagRepository.GetArticleTagList(article.ID)
	if err != nil {
		return nil, err
	}

	// 查询推荐文章
	rmArticle, err := s.svcCtx.ArticleRepository.GetRecommendArticle(article.CategoryID)
	if err != nil {
		return nil, err
	}
	// 查询最新文章
	page := &request.PageInfo{
		Page:     0,
		PageSize: 5,
		Order:    "id",
		OrderKey: "desc",
	}
	newestArticle, _, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, err
	}
	// 查询上一篇文章
	lastArticle, err := s.svcCtx.ArticleRepository.GetLastArticle(id)
	if err != nil {
		return nil, err
	}
	// 查询下一篇文章
	nextArticle, err := s.svcCtx.ArticleRepository.GetNextArticle(id)
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
func (s *ArticleService) GetArticleList(reqCtx *request.Context, page *request.PageInfo) (list []*response.ArticleDTO, total int64, err error) {
	// 查询文章列表
	articles, total, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.GetCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.GetArticleTagList(article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}
	return list, total, err
}

func (s *ArticleService) GetArticleListByCondition(reqCtx *request.Context, req *request.ArticleCondition) (data *response.ArticleConditionDTO, err error) {
	resp := &response.ArticleConditionDTO{}

	// 查询文章列表
	var articles []*entity.Article

	if req.CategoryID != 0 {
		category, err := s.svcCtx.CategoryRepository.GetCategory(reqCtx, req.CategoryID)
		if err != nil {
			return nil, err
		}
		articles, _, err = s.svcCtx.ArticleRepository.GetArticleListByCategoryId(category.ID)
		resp.ConditionName = category.CategoryName
	} else if req.TagID != 0 {
		tag, err := s.svcCtx.TagRepository.GetTag(reqCtx, req.TagID)
		if err != nil {
			return nil, err
		}
		articles, _, err = s.svcCtx.ArticleRepository.GetArticleListByTagId(tag.ID)
		resp.ConditionName = tag.TagName
	}

	var list []*response.ArticleDTO
	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.GetCategory(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.GetArticleTagList(article.ID)

		articleVO := convertArticle(article)
		articleVO.CategoryName = category.CategoryName
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}

	resp.ArticleDTOList = list
	return resp, err
}

func (s *ArticleService) GetArticleArchives(reqCtx *request.Context, page *request.PageInfo) (list []*response.ArticleRecommendDTO, total int64, err error) {
	page.Order = "id"
	page.OrderKey = "desc"
	newestArticle, total, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	return convertRecommendArticles(newestArticle), total, err
}

func convertResponseArticle(entity *entity.Article) *response.ArticleDetails {
	return &response.ArticleDetails{
		ID:             entity.ID,
		ArticleCover:   entity.ArticleCover,
		ArticleTitle:   entity.ArticleTitle,
		ArticleContent: entity.ArticleContent,
		//LikeCount:            entity.LikeCount,
		//ViewsCount:           entity.ViewsCount,
		Type:        entity.Type,
		OriginalURL: entity.OriginalUrl,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		CategoryID:  entity.CategoryID,
		//CategoryName:         "",
		ArticleTagList: []*response.TagDTO{{1, "tag1"}, {2, "tag2"}, {3, "tag3"}},
		//LastArticle:          response.ArticlePaginationDTO{},
		//NextArticle:          response.ArticlePaginationDTO{},
		//RecommendArticleList: nil,
		//NewestArticleList:    nil,
	}
}

func convertRecommendArticles(list []*entity.Article) []*response.ArticleRecommendDTO {
	var out []*response.ArticleRecommendDTO
	for _, item := range list {
		at := &response.ArticleRecommendDTO{
			ID:           item.ID,
			ArticleCover: item.ArticleCover,
			ArticleTitle: item.ArticleTitle,
			CreatedAt:    item.CreatedAt,
		}
		out = append(out, at)
	}

	return out
}

func convertArticlePagination(article *entity.Article) *response.ArticlePaginationDTO {
	if article == nil {
		return nil
	}
	return &response.ArticlePaginationDTO{
		ID:           article.ID,
		ArticleCover: article.ArticleCover,
		ArticleTitle: article.ArticleTitle,
	}
}

func convertTagList(tags []*entity.Tag) []*response.TagDTO {
	var tagList []*response.TagDTO
	for _, tag := range tags {
		tagList = append(tagList, &response.TagDTO{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	return tagList
}

func convertArticle(article *entity.Article) *response.ArticleDTO {
	out := &response.ArticleDTO{
		ID:             article.ID,
		ArticleCover:   article.ArticleCover,
		ArticleTitle:   article.ArticleTitle,
		ArticleContent: article.ArticleContent,
		LikeCount:      100,
		ViewsCount:     200,
		Type:           article.Type,
		OriginalURL:    article.OriginalUrl,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		CategoryID:     article.CategoryID,
	}
	return out
}
