package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

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
