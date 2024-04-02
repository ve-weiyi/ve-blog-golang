package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCleanMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanMenuListLogic {
	return &CleanMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanMenuListLogic) CleanMenuList(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.BatchResp, err error) {
	in := convert.EmptyReq()
	out, err := l.svcCtx.MenuRpc.CleanMenuList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, err
}
