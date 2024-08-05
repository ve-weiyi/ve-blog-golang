package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertMenuTypes(in *blogrpc.Menu) (out *types.MenuDetails) {
	jsonconv.ObjectToObject(in, &out)

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}

func ConvertMenuDetailsTypes(in *blogrpc.MenuDetails) (out *types.MenuDetails) {
	out = &types.MenuDetails{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Title,
		Type:      in.Type,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Meta:      types.Meta{},
		Children:  make([]*types.MenuDetails, 0),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	jsonconv.JsonToObject(in.Extra, &out.Meta)

	for _, v := range in.Children {
		out.Children = append(out.Children, ConvertMenuDetailsTypes(v))
	}

	return
}

func ConvertMenuPb(in *types.MenuDetails) (out *blogrpc.Menu) {
	jsonconv.ObjectToObject(in, &out)

	out.Extra = jsonconv.ObjectToJson(in.Meta)
	return
}

func ConvertMenuDetailPb(in types.RouteConfigsTable) (out *blogrpc.MenuDetails) {
	out = &blogrpc.MenuDetails{
		Id:        0,
		ParentId:  0,
		Title:     "",
		Type:      0,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Extra:     jsonconv.ObjectToJson(in.Meta),
		CreatedAt: 0,
		UpdatedAt: 0,
		Children:  make([]*blogrpc.MenuDetails, 0),
	}

	for _, v := range in.Children {
		out.Children = append(out.Children, ConvertMenuDetailPb(v))
	}

	return
}

func ConvertSyncMenuPb(in *types.SyncMenuReq) (out *blogrpc.SyncMenuReq) {
	list := make([]*blogrpc.MenuDetails, 0)
	for _, m := range in.Menus {
		list = append(list, ConvertMenuDetailPb(m))
	}

	out = &blogrpc.SyncMenuReq{
		Menus: list,
	}

	return
}
