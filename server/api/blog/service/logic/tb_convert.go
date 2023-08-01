package logic

import (
	"math/rand"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
)

func convertAuthConfig(conf properties.AuthConfig) *oauth.AuthConfig {
	return &oauth.AuthConfig{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectUrl:  conf.RedirectUrl,
	}
}

func convertTagList(list []*entity.Tag) []*response.TagDTO {
	var tagList []*response.TagDTO
	for _, tag := range list {
		tagList = append(tagList, &response.TagDTO{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	return tagList
}

func convertCategoryList(list []*entity.Category) []*response.CategoryDTO {
	var categoryList []*response.CategoryDTO

	for _, in := range list {
		data := &response.CategoryDTO{
			ID:           in.ID,
			CategoryName: in.CategoryName,
			ArticleCount: 0,
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}
		categoryList = append(categoryList, data)
	}
	return categoryList
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

func convertArticleStatisticsList(list []*entity.Article) []*response.ArticleStatisticsDTO {
	var dateMap = make(map[string]int64)
	for _, item := range list {
		key := item.CreatedAt.Format("2006-01-02")
		if _, ok := dateMap[key]; ok {
			dateMap[key]++
		} else {
			dateMap[key] = 1
		}
	}
	// 日期统计
	var out []*response.ArticleStatisticsDTO
	for k, v := range dateMap {
		at := &response.ArticleStatisticsDTO{
			Day:   k,
			Count: v,
		}
		out = append(out, at)
	}

	return out
}

func convertUniqueViewList(list []*entity.UniqueView) []*response.UniqueViewDTO {
	var out []*response.UniqueViewDTO
	for _, item := range list {
		at := &response.UniqueViewDTO{
			Day:   item.CreatedAt.Format("2006-01-02"),
			Count: rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}

func convertArticleRankList(list []*entity.Article) []*response.ArticleRankDTO {
	var out []*response.ArticleRankDTO
	for _, item := range list {
		at := &response.ArticleRankDTO{
			ID:           item.ID,
			ArticleTitle: item.ArticleTitle,
			Count:        rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}
