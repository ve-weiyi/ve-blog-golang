package auth

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.EmptyReq) (resp *types.EmptyResp, err error) {
	uid := l.ctx.Value("uid").(string)

	in := &accountrpc.LogoutReq{
		UserId: cast.ToInt64(l.ctx.Value("uid")),
	}

	out, err := l.svcCtx.AccountRpc.Logout(l.ctx, in)
	if err != nil {
		return nil, err
	}

	redisKey := middlewarex.GetUserLogoutKey(cast.ToInt64(uid))
	_ = l.svcCtx.Redis.SetexCtx(l.ctx, redisKey, fmt.Sprintf("%d", out.LogoutAt), 7*24*60*60)

	return &types.EmptyResp{}, nil
}
