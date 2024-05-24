package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendForgetEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送忘记密码邮件
func NewSendForgetEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendForgetEmailLogic {
	return &SendForgetEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendForgetEmailLogic) SendForgetEmail(reqCtx *types.RestHeader, req *types.UserEmailReq) (resp *types.EmptyResp, err error) {
	in := &blog.UserEmailReq{
		Username: req.Username,
	}

	_, err = l.svcCtx.AuthRpc.SendResetPasswordEmail(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
