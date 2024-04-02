package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgetPasswordEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordEmailLogic {
	return &ForgetPasswordEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送忘记密码邮件
func (l *ForgetPasswordEmailLogic) ForgetPasswordEmail(in *account.UserEmailReq) (*account.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &account.EmptyResp{}, nil
}
