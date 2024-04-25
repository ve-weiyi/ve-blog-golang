package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/menurpc"

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

func (l *SyncMenuListLogic) SyncMenuList(reqCtx *types.RestHeader, req *types.SyncMenuRequest) (resp *types.BatchResp, err error) {
	in := &menurpc.SyncMenuRequest{}
	err = jsonconv.ObjectMarshal(req, in)
	if err != nil {
		return nil, err
	}

	out, err := l.svcCtx.MenuRpc.SyncMenuList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, err
}
