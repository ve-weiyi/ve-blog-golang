package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendResetEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送重置密码邮件
func NewSendResetEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendResetEmailLogic {
	return &SendResetEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendResetEmailLogic) SendResetEmail(req *types.UserEmailReq) (resp *types.EmptyResp, err error) {
	in := &blog.UserEmailReq{
		Username: req.Username,
	}

	_, err = l.svcCtx.AuthRpc.SendResetPasswordEmail(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
