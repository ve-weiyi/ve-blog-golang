package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertTagPb(in *types.Tag) (out *blogrpc.Tag) {
	return &blogrpc.Tag{
		Id:        in.Id,
		TagName:   in.TagName,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func ConvertTagTypes(in *blogrpc.Tag) (out *types.Tag) {
	return &types.Tag{
		Id:        in.Id,
		TagName:   in.TagName,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func ConvertTagDetailsTypes(in *blogrpc.Tag) (out *types.TagDetails) {
	return &types.TagDetails{
		Id:           in.Id,
		TagName:      in.TagName,
		ArticleCount: 0,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
