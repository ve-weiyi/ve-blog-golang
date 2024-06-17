package permissionrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncApiListLogic {
	return &SyncApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步接口列表
func (l *SyncApiListLogic) SyncApiList(in *permissionrpc.EmptyReq) (*permissionrpc.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.BatchResp{}, nil
}
