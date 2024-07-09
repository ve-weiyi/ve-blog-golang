package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertArticleHomeTypes(in *blog.Article) (out *types.ArticleHomeDTO) {
	out = &types.ArticleHomeDTO{
		Id:             in.Id,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		Type:           in.Type,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		CategoryName:   "",
		TagNameList:    make([]string, 0),
		LikeCount:      0,
		ViewsCount:     0,
	}

	return
}

func ConvertArticlePreviewTypes(in *blog.Article) (out *types.ArticlePreviewDTO) {
	out = &types.ArticlePreviewDTO{
		Id:           in.Id,
		ArticleCover: in.ArticleCover,
		ArticleTitle: in.ArticleTitle,
		CreatedAt:    in.CreatedAt,
	}
	return
}
