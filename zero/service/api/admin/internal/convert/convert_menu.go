package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertMenuTypes(in *blog.Menu) (out *types.MenuDetails) {
	jsonconv.ObjectToObject(in, &out)

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}

func ConvertMenuDetailsTypes(in *blog.MenuDetails) (out *types.MenuDetails) {
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

func ConvertMenuPb(in *types.MenuDetails) (out *blog.Menu) {
	jsonconv.ObjectToObject(in, &out)

	out.Extra = jsonconv.ObjectToJson(in.Meta)
	return
}

func ConvertMenuDetailPb(in types.RouteConfigsTable) (out *blog.MenuDetails) {
	out = &blog.MenuDetails{
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
		Children:  make([]*blog.MenuDetails, 0),
	}

	for _, v := range in.Children {
		out.Children = append(out.Children, ConvertMenuDetailPb(v))
	}

	return
}

func ConvertSyncMenuPb(in *types.SyncMenuReq) (out *blog.SyncMenuReq) {
	list := make([]*blog.MenuDetails, 0)
	for _, m := range in.Menus {
		list = append(list, ConvertMenuDetailPb(m))
	}

	out = &blog.SyncMenuReq{
		Menus: list,
	}

	return
}
