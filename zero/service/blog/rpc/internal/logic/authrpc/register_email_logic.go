package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
func (l *RegisterEmailLogic) RegisterEmail(in *blog.UserEmailReq) (*blog.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &blog.EmptyResp{}, nil
}
