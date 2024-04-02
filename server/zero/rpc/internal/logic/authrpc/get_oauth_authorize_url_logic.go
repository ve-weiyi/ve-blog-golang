package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthAuthorizeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthAuthorizeUrlLogic {
	return &GetOauthAuthorizeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取授权地址
func (l *GetOauthAuthorizeUrlLogic) GetOauthAuthorizeUrl(in *account.OauthLoginReq) (*account.OauthLoginUrlResp, error) {
	// todo: add your logic here and delete this line

	return &account.OauthLoginUrlResp{}, nil
}
