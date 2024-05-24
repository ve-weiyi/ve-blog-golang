package auth

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/authrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注销
func NewLogoffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoffLogic {
	return &LogoffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoffLogic) Logoff(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.EmptyResp, err error) {
	in := authrpc.LogoffReq{
		UserId: cast.ToInt64(reqCtx.HeaderXUserId),
	}

	_, err = l.svcCtx.AuthRpc.Logoff(l.ctx, &in)
	if err != nil {
		return
	}

	return &types.EmptyResp{}, nil
}
