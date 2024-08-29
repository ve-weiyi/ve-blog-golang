package permissionrpclogic

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
)

func convertApiIn(in *permissionrpc.ApiNewReq) (out *model.Api) {
	out = &model.Api{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
	}

	return out
}

func convertApiOut(in *model.Api) (out *permissionrpc.ApiDetails) {
	out = &permissionrpc.ApiDetails{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
