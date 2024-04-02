package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthAuthorizeUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOauthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthAuthorizeUrlLogic {
	return &GetOauthAuthorizeUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOauthAuthorizeUrlLogic) GetOauthAuthorizeUrl(req *types.OauthLoginReq) (resp *types.OauthLoginUrl, err error) {
	// todo: add your logic here and delete this line

	return
}
