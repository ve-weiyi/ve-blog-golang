package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthAuthorizeUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 第三方登录授权地址
func NewOauthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthAuthorizeUrlLogic {
	return &OauthAuthorizeUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthAuthorizeUrlLogic) OauthAuthorizeUrl(reqCtx *types.RestHeader, req *types.OauthLoginReq) (resp *types.OauthLoginUrlResp, err error) {
	// todo: add your logic here and delete this line

	return
}
