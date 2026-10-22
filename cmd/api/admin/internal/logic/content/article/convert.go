package article

import (
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/contentservice"
)

func convertArticleVO(article *contentservice.Article) *types.ArticleVO {
	if article == nil {
		return nil
	}

	tagNames := make([]string, 0, len(article.Tags))
	for _, tag := range article.Tags {
		tagNames = append(tagNames, tag.TagName)
	}

	categoryName := ""
	if article.Category != nil {
		categoryName = article.Category.CategoryName
	}

	return &types.ArticleVO{
		Id:             article.Id,
		ArticleCover:   article.ArticleCover,
		ArticleTitle:   article.ArticleTitle,
		ArticleContent: article.ArticleContent,
		ArticleType:    article.ArticleType,
		OriginalUrl:    article.OriginalUrl,
		IsTop:          article.IsTop,
		IsDelete:       article.IsDelete,
		Status:         article.Status,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		CategoryName:   categoryName,
		TagNameList:    tagNames,
		LikeCount:      article.LikeCount,
		ViewCount:      article.ViewCount,
	}
}
