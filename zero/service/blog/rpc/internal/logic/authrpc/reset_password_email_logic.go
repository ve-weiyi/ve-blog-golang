package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordEmailLogic {
	return &ResetPasswordEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送忘记密码邮件
func (l *ResetPasswordEmailLogic) ResetPasswordEmail(in *blog.UserEmailReq) (*blog.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &blog.EmptyResp{}, nil
}
