package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/accountrpc"
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

	in := accountrpc.LoginReq{
		Username: req.Username,
		Password: req.Password,
		Code:     req.Code,
	}

	out, err := l.svcCtx.AccountRpc.Login(l.ctx, &in)
	if err != nil {
		return
	}

	tk, err := l.createToken(out.UserId, out.Username, "")
	if err != nil {
		return
	}

	resp = &types.LoginResp{
		Token:    tk,
		UserInfo: convertUserInfo(out),
	}

	return
}

func (l *LoginLogic) createToken(uid int64, username string, loginType string) (token *types.Token, err error) {
	now := time.Now().Unix()
	expiresIn := time.Now().Add(7 * 24 * time.Hour).Unix()
	refreshExpiresIn := time.Now().Add(30 * 24 * time.Hour).Unix()
	issuer := "blog"

	accessToken, err := l.svcCtx.Token.CreateToken(
		jjwt.TokenExt{
			Uid:       int(uid),
			Username:  username,
			LoginType: loginType,
		},
		jwt.StandardClaims{
			ExpiresAt: expiresIn,
			IssuedAt:  now,
			Issuer:    issuer,
		})

	refreshToken, err := l.svcCtx.Token.CreateToken(
		jjwt.TokenExt{
			Uid:       int(uid),
			Username:  username,
			LoginType: loginType,
		},
		jwt.StandardClaims{
			ExpiresAt: refreshExpiresIn,
			IssuedAt:  now,
			Issuer:    issuer,
		})

	token = &types.Token{
		TokenType:        "Bearer",
		AccessToken:      accessToken,
		ExpiresIn:        expiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
	}

	//生成token
	return token, nil
}

func convertUserInfo(in *accountrpc.LoginResp) (out *types.UserInfo) {
	out = &types.UserInfo{
		UserId:   in.UserId,
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Intro:    in.Intro,
		Website:  in.Website,
		Email:    in.Email,
	}
	return
}
