package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertTagTypes(in *blogrpc.Tag) (out *types.Tag) {
	return &types.Tag{
		Id:           in.Id,
		TagName:      in.TagName,
		ArticleCount: 0,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
