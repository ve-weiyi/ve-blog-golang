package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方登录
func (l *OauthLoginLogic) OauthLogin(in *account.OauthLoginReq) (*account.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &account.LoginResp{}, nil
}
