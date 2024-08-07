package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建api路由
func NewAddApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddApiLogic {
	return &AddApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddApiLogic) AddApi(req *types.ApiNew) (resp *types.ApiBackDTO, err error) {
	in := &permissionrpc.ApiNew{
		Id:        0,
		Name:      req.Name,
		Path:      req.Path,
		Method:    req.Method,
		ParentId:  req.ParentId,
		Traceable: req.Traceable,
		Status:    req.Status,
	}
	_, err = l.svcCtx.PermissionRpc.AddApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return
}
