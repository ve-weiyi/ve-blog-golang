package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/menurpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMenuListLogic {
	return &SyncMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncMenuListLogic) SyncMenuList(req *types.SyncMenuRequest) (resp *types.BatchResult, err error) {
	in := &menurpc.SyncMenuRequest{}
	err = jsonconv.ObjectMarshal(req, in)
	if err != nil {
		return nil, err
	}

	out, err := l.svcCtx.MenuRpc.SyncMenuList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResult{
		SuccessCount: out.SuccessCount,
	}
	return resp, err
}
