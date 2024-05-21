package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendRegisterEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送注册账号邮件
func NewSendRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendRegisterEmailLogic {
	return &SendRegisterEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendRegisterEmailLogic) SendRegisterEmail(reqCtx *types.RestHeader, req *types.UserEmailReq) (resp *types.EmptyResp, err error) {
	in := &blog.UserEmailReq{
		Username: req.Username,
	}

	_, err = l.svcCtx.AuthRpc.RegisterEmail(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
