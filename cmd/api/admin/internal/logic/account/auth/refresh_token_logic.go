package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新token
func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.LoginResp, err error) {
	deviceId, _ := metax.GetApiDeviceIdFromCtx(l.ctx)
	tk, err := l.svcCtx.TokenManager.Refresh(l.ctx, tokenx.RefreshRequest{
		RefreshToken: req.RefreshToken,
		UserID:       req.UserId,
		DeviceID:     deviceId,
	})
	if err != nil {
		return nil, toAuthError(err)
	}

	return &types.LoginResp{
		UserId: tk.UserID,
		Scope:  l.svcCtx.Config.Name,
		Token:  toToken(tk),
	}, nil
}
