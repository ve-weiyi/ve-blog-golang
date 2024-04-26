package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertArticleTypes(in *blog.Article) (out *types.Article) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertArticlePb(in *types.Article) (out *blog.Article) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertArticleDetailsTypes(in *blog.Article) (out *types.ArticleDetailsResp) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertArticleHomeTypes(in *blog.Article) (out *types.ArticleHome) {
	out = &types.ArticleHome{
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
