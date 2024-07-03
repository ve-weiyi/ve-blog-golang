package auth

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/authrpc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	in := authrpc.LoginReq{
		Username: req.Username,
		Password: req.Password,
		Code:     req.Code,
	}

	out, err := l.svcCtx.AuthRpc.Login(l.ctx, &in)
	if err != nil {
		return
	}

	tk, err := l.createToken(out.UserId, out.Username, "")
	if err != nil {
		return
	}

	resp = &types.LoginResp{
		Token:    tk,
		UserInfo: convert.ConvertUserInfo(out),
	}

	return
}

func (l *LoginLogic) createToken(uid int64, username string, loginType string) (token *types.Token, err error) {
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

	//生成token
	return token, nil
}
