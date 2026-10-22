package auth

import (
	"context"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/authservice"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

func onLogin(ctx context.Context, svcCtx *svc.ServiceContext, login *authservice.LoginResponse) (resp *types.LoginResp, err error) {
	deviceId, _ := metax.GetApiDeviceIdFromCtx(ctx)
	tk, err := svcCtx.TokenManager.Issue(ctx, tokenx.IssueRequest{
		UserID:   login.UserId,
		DeviceID: deviceId,
		Issuer:   svcCtx.Config.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		UserId: login.UserId,
		Scope:  svcCtx.Config.Name,
		Token:  toToken(tk),
	}, nil
}
