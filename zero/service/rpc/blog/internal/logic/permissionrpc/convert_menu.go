package permissionrpclogic

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
)

func convertMenuIn(in *permissionrpc.MenuNewReq) (out *model.TMenu) {
	out = &model.TMenu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Name,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Type:      in.Type,
		Extra:     in.Extra,
	}

	return out
}

func convertMenuOut(in *model.TMenu) (out *permissionrpc.MenuDetails) {
	out = &permissionrpc.MenuDetails{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Title:     in.Name,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Type:      in.Type,
		Extra:     in.Extra,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}
	return out
}
