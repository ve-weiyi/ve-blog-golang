package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterEmailLogic {
	return &RegisterEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送注册邮件
func (l *RegisterEmailLogic) RegisterEmail(in *account.UserEmailReq) (*account.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &account.EmptyResp{}, nil
}
