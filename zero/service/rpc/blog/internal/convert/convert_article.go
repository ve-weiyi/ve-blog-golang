package convert

import (
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
		Type:           in.Type,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		IsDelete:       in.IsDelete,
		Status:         in.Status,
		// CreatedAt: time.Unix(in.CreatedAt, 0),
		// UpdatedAt: time.Unix(in.UpdatedAt, 0),
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
		Type:           in.Type,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		IsDelete:       in.IsDelete,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}
