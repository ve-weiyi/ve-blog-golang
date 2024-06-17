package convert

import (
	"math/rand"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

func ConvertHomeTagTypes(in *blogrpc.Tag) (out *types.TagDTO) {
	return &types.TagDTO{
		Id:      in.Id,
		TagName: in.TagName,
	}
}

func ConvertHomeCategoryTypes(in *blogrpc.Category) (out *types.CategoryDTO) {
	return &types.CategoryDTO{
		Id:           in.Id,
		CategoryName: in.CategoryName,
	}
}

func ConvertHomeArticleRankTypes(in *blogrpc.Article) (out *types.ArticleViewRankDTO) {
	return &types.ArticleViewRankDTO{
		Id:           in.Id,
		ArticleTitle: in.ArticleTitle,
		Count:        rand.Int63n(100),
	}
}

func ConvertHomeViewTypes(in *blogrpc.UserVisit) (out *types.UniqueViewDTO) {

	return &types.UniqueViewDTO{
		Date:  in.Date,
		Count: in.ViewCount,
	}
}
