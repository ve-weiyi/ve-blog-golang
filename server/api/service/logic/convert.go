package logic

import (
	"math/rand"

	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	response2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
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

func convertTagList(list []*entity2.Tag) []*response2.TagDTO {
	var tagList []*response2.TagDTO
	for _, tag := range list {
		tagList = append(tagList, &response2.TagDTO{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	return tagList
}

func convertCategoryList(list []*entity2.Category) []*response2.CategoryDTO {
	var categoryList []*response2.CategoryDTO

	for _, in := range list {
		data := &response2.CategoryDTO{
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

func convertResponseArticle(entity *entity2.Article) *response2.ArticleDetails {
	return &response2.ArticleDetails{
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
		ArticleTagList: []*response2.TagDTO{{1, "tag1"}, {2, "tag2"}, {3, "tag3"}},
		//LastArticle:          response.ArticlePaginationDTO{},
		//NextArticle:          response.ArticlePaginationDTO{},
		//RecommendArticleList: nil,
		//NewestArticleList:    nil,
	}
}

func convertRecommendArticles(list []*entity2.Article) []*response2.ArticleRecommendDTO {
	var out []*response2.ArticleRecommendDTO
	for _, item := range list {
		at := &response2.ArticleRecommendDTO{
			ID:           item.ID,
			ArticleCover: item.ArticleCover,
			ArticleTitle: item.ArticleTitle,
			CreatedAt:    item.CreatedAt,
		}
		out = append(out, at)
	}

	return out
}

func convertArticlePagination(article *entity2.Article) *response2.ArticlePaginationDTO {
	if article == nil {
		return nil
	}
	return &response2.ArticlePaginationDTO{
		ID:           article.ID,
		ArticleCover: article.ArticleCover,
		ArticleTitle: article.ArticleTitle,
	}
}

func convertArticle(article *entity2.Article) *response2.ArticleDTO {
	out := &response2.ArticleDTO{
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

func convertArticleStatisticsList(list []*entity2.Article) []*response2.ArticleStatisticsDTO {
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
	var out []*response2.ArticleStatisticsDTO
	for k, v := range dateMap {
		at := &response2.ArticleStatisticsDTO{
			Day:   k,
			Count: v,
		}
		out = append(out, at)
	}

	return out
}

func convertUniqueViewList(list []*entity2.UniqueView) []*response2.UniqueViewDTO {
	var out []*response2.UniqueViewDTO
	for _, item := range list {
		at := &response2.UniqueViewDTO{
			Day:   item.CreatedAt.Format("2006-01-02"),
			Count: rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}

func convertArticleRankList(list []*entity2.Article) []*response2.ArticleRankDTO {
	var out []*response2.ArticleRankDTO
	for _, item := range list {
		at := &response2.ArticleRankDTO{
			ID:           item.ID,
			ArticleTitle: item.ArticleTitle,
			Count:        rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}
