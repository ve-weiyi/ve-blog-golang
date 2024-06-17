package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 同步api列表
func NewSyncApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncApiListLogic {
	return &SyncApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncApiListLogic) SyncApiList(req *types.EmptyReq) (resp *types.BatchResp, err error) {
	in := &permissionrpc.EmptyReq{}

	out, err := l.svcCtx.PermissionRpc.SyncApiList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
