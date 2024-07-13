package menu

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
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

func (l *SyncMenuListLogic) SyncMenuList(req *types.SyncMenuReq) (resp *types.BatchResp, err error) {

	in := convert.ConvertSyncMenuPb(req)

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
