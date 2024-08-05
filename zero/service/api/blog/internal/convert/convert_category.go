package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertCategoryTypes(in *blogrpc.Category) (out *types.Category) {
	return &types.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
