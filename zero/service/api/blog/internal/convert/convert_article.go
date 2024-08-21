package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertArticleHomeTypes(in *blogrpc.ArticleDetails) (out *types.ArticleHome) {
	out = &types.ArticleHome{
		Id:             in.Id,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		CategoryName:   "",
		TagNameList:    make([]string, 0),
		LikeCount:      in.LikeCount,
		ViewsCount:     0,
	}

	return
}

func ConvertArticlePreviewTypes(in *blogrpc.ArticleDetails) (out *types.ArticlePreview) {
	out = &types.ArticlePreview{
		Id:           in.Id,
		ArticleCover: in.ArticleCover,
		ArticleTitle: in.ArticleTitle,
		CreatedAt:    in.CreatedAt,
	}
	return
}
