package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertCategoryPb(in *types.Category) (out *blogrpc.Category) {
	return &blogrpc.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func ConvertCategoryTypes(in *blogrpc.Category) (out *types.Category) {
	return &types.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func ConvertCategoryDetailsTypes(in *blogrpc.Category) (out *types.CategoryDetails) {
	return &types.CategoryDetails{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		ArticleCount: 0,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
