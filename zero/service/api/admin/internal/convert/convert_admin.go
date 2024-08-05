package convert

import (
	"math/rand"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

func ConvertHomeTagTypes(in *blogrpc.Tag) (out *types.TagDTO) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertHomeCategoryTypes(in *blogrpc.Category) (out *types.CategoryDTO) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertHomeArticleStaticsTypes(in *blogrpc.Article) (out *types.ArticleStatisticsDTO) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertHomeArticleRankTypes(in *blogrpc.Article) (out *types.ArticleViewRankDTO) {
	jsonconv.ObjectToObject(in, &out)

	out.Count = rand.Int63n(100)
	return
}

func ConvertHomeViewTypes(in *blogrpc.UserVisit) (out *types.UniqueViewDTO) {

	return &types.UniqueViewDTO{
		Date:  in.Date,
		Count: in.ViewCount,
	}
}
