package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新api路由
func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiLogic) UpdateApi(req *types.ApiNewReq) (resp *types.ApiBackDTO, err error) {
	in := &permissionrpc.ApiNewReq{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Method:    req.Method,
		Traceable: req.Traceable,
		IsDisable: req.IsDisable,
		Children:  nil,
	}

	out, err := l.svcCtx.PermissionRpc.UpdateApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertApiTypes(out), nil
}
