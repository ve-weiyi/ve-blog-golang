package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
	in := &blog.OauthLoginReq{
		Platform: req.Platform,
		Code:     req.Code,
		State:    req.State,
	}

	out, err := l.svcCtx.AuthRpc.GetOauthAuthorizeUrl(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.OauthLoginUrlResp{
		Url: out.Url,
	}
	return resp, nil
}
