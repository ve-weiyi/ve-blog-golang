package auth

import (
	"context"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
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
	expires := 7 * 24 * time.Hour
	uid := login.UserId

	accessToken, err := svcCtx.TokenHolder.CreateToken(
		ctx,
		cast.ToString(uid),
		expires,
	)

	token = &types.Token{
		UserId:      uid,
		TokenType:   "Bearer",
		AccessToken: accessToken,
		ExpiresIn:   time.Now().Add(expires).Unix(),
	}

	// 生成token
	return token, nil
}
