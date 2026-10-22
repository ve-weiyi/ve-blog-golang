package authservicelogic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/authrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/auth/oauthstate"
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
func (l *GetOAuthAuthorizeUrlLogic) GetOAuthAuthorizeUrl(in *authrpc.GetOAuthAuthorizeUrlRequest) (*authrpc.GetOAuthAuthorizeUrlResponse, error) {
	auth, ok := l.svcCtx.OAuthProviders[in.Platform]
	if !ok {
		return nil, fmt.Errorf("platform %s is not support", in.Platform)
	}

	nonce, err := oauthstate.Issue(l.ctx, l.svcCtx.Redis, in.Platform, in.State)
	if err != nil {
		return nil, err
	}

	return &authrpc.GetOAuthAuthorizeUrlResponse{
		AuthorizeUrl: auth.GetAuthLoginUrl(nonce),
		State:        nonce,
	}, nil
}
