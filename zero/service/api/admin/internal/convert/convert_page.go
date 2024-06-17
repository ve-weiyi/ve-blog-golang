package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertPagePb(in *types.Page) (out *blogrpc.Page) {
	return &blogrpc.Page{
		Id:        in.Id,
		PageName:  in.PageName,
		PageLabel: in.PageLabel,
		PageCover: in.PageCover,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func ConvertPageTypes(in *blogrpc.Page) (out *types.Page) {
	return &types.Page{
		Id:        in.Id,
		PageName:  in.PageName,
		PageLabel: in.PageLabel,
		PageCover: in.PageCover,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
