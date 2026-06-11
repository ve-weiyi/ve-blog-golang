package userauthservicelogic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/userauthrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type GetOAuthAuthorizeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOAuthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAuthAuthorizeUrlLogic {
	return &GetOAuthAuthorizeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取第三方OAuth授权地址
func (l *GetOAuthAuthorizeUrlLogic) GetOAuthAuthorizeUrl(in *userauthrpc.GetOAuthAuthorizeUrlRequest) (*userauthrpc.GetOAuthAuthorizeUrlResponse, error) {
	auth, ok := l.svcCtx.OAuthProviders[in.Platform]
	if !ok {
		return nil, fmt.Errorf("platform %s is not support", in.Platform)
	}

	return &userauthrpc.GetOAuthAuthorizeUrlResponse{
		AuthorizeUrl: auth.GetAuthLoginUrl(in.State),
	}, nil
}
