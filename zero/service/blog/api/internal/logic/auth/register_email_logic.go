package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterEmailLogic {
	return &RegisterEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterEmailLogic) RegisterEmail(reqCtx *types.RestHeader, req *types.UserEmailReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
