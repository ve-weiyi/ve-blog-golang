package authrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
func (l *GetOauthAuthorizeUrlLogic) GetOauthAuthorizeUrl(in *blog.OauthLoginReq) (*blog.OauthLoginUrlResp, error) {
	var auth oauth.Oauth
	for platform, v := range l.svcCtx.Oauth {
		if platform == in.Platform {
			auth = v
		}
	}

	if auth == nil {
		return nil, fmt.Errorf("platform %s is not support", in.Platform)
	}

	resp := &blog.OauthLoginUrlResp{}
	resp.Url = auth.GetRedirectUrl(in.State)
	return resp, nil
}
