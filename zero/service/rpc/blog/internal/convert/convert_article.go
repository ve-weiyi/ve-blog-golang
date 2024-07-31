package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertArticlePbToModel(in *blog.Article) (out *model.Article) {
	out = &model.Article{
		Id:             in.Id,
		UserId:         in.UserId,
		CategoryId:     in.CategoryId,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		IsDelete:       in.IsDelete,
		Status:         in.Status,
		LikeCount:      in.LikeCount,
		CreatedAt:      time.Unix(in.CreatedAt, 0),
		UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertArticleModelToPb(in *model.Article) (out *blog.Article) {
	out = &blog.Article{
		Id:             in.Id,
		UserId:         in.UserId,
		CategoryId:     in.CategoryId,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		IsDelete:       in.IsDelete,
		Status:         in.Status,
		LikeCount:      in.LikeCount,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}
