package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/userauthservice"
)

func onLogin(ctx context.Context, svcCtx *svc.ServiceContext, login *userauthservice.LoginResponse) (resp *types.LoginResp, err error) {
	tk, err := svcCtx.TokenStore.GenerateToken(login.UserId)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		UserId: login.UserId,
		Scope:  svcCtx.Config.Name,
		Token: &types.Token{
			TokenType:        tk.TokenType,
			AccessToken:      tk.AccessToken,
			ExpiresIn:        tk.ExpiresIn,
			RefreshToken:     tk.RefreshToken,
			RefreshExpiresIn: tk.RefreshExpiresIn,
			RefreshExpiresAt: tk.RefreshExpiresAt,
		},
	}, nil
}
