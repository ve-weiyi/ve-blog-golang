package auth

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 第三方登录
func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.LoginResp, err error) {
	in := accountrpc.OauthLoginReq{
		Platform: req.Platform,
		Code:     req.Code,
		State:    req.State,
	}

	out, err := l.svcCtx.AccountRpc.OauthLogin(l.ctx, &in)
	if err != nil {
		return
	}

	tk, err := l.createToken(out.UserId, out.Username, req.Platform)
	if err != nil {
		return
	}

	resp = &types.LoginResp{
		Token: tk,
	}
	return
}

func (l *OauthLoginLogic) createToken(uid int64, username string, loginType string) (token *types.Token, err error) {
	now := time.Now().Unix()
	expiresIn := time.Now().Add(7 * 24 * time.Hour).Unix()
	refreshExpiresIn := time.Now().Add(30 * 24 * time.Hour).Unix()
	issuer := "blog"

	accessToken, err := l.svcCtx.Token.CreateToken(
		jtoken.WithExpiresAt(expiresIn),
		jtoken.WithIssuedAt(now),
		jtoken.WithIssuer(issuer),
		jtoken.WithClaimExt("uid", uid),
		jtoken.WithClaimExt("username", username),
		jtoken.WithClaimExt("login_type", loginType),
	)

	refreshToken, err := l.svcCtx.Token.CreateToken(
		jtoken.WithExpiresAt(refreshExpiresIn),
		jtoken.WithIssuedAt(now),
		jtoken.WithIssuer(issuer),
		jtoken.WithClaimExt("uid", uid),
		jtoken.WithClaimExt("username", username),
		jtoken.WithClaimExt("login_type", loginType),
	)

	token = &types.Token{
		UserId:           uid,
		TokenType:        "Bearer",
		AccessToken:      accessToken,
		ExpiresIn:        expiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
	}

	// 生成token
	return token, nil
}
