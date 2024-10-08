package auth

import (
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
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
		Username:   req.Username,
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	}

	out, err := l.svcCtx.AccountRpc.Login(l.ctx, &in)
	if err != nil {
		return
	}

	tk, err := createToken(l.ctx, l.svcCtx, out)
	if err != nil {
		return
	}

	resp = &types.LoginResp{
		Token: tk,
	}

	return
}

func createToken(ctx context.Context, svcCtx *svc.ServiceContext, login *accountrpc.LoginResp) (token *types.Token, err error) {
	now := time.Now().Unix()
	expiresIn := time.Now().Add(7 * 24 * time.Hour).Unix()
	refreshExpiresIn := time.Now().Add(30 * 24 * time.Hour).Unix()
	issuer := "blog"
	uid := login.UserId
	var roles []string
	for _, role := range login.Roles {
		roles = append(roles, role.RoleName)
	}

	accessToken, err := svcCtx.Token.CreateToken(
		jtoken.WithExpiresAt(expiresIn),
		jtoken.WithIssuedAt(now),
		jtoken.WithIssuer(issuer),
		jtoken.WithClaimExt("uid", uid),
		jtoken.WithClaimExt("roles", strings.Join(roles, ",")),
	)

	refreshToken, err := svcCtx.Token.CreateToken(
		jtoken.WithExpiresAt(refreshExpiresIn),
		jtoken.WithIssuedAt(now),
		jtoken.WithIssuer(issuer),
		jtoken.WithClaimExt("uid", uid),
		jtoken.WithClaimExt("roles", strings.Join(roles, ",")),
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
