package auth

import (
	"context"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/accessservice"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/authservice"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

func onLogin(ctx context.Context, svcCtx *svc.ServiceContext, login *authservice.LoginResponse) (resp *types.LoginResp, err error) {
	// 管理端只允许拥有启用角色的账号登录
	roles, err := svcCtx.AccessService.ListUserRoles(ctx, &accessservice.ListUserRolesRequest{UserId: login.UserId})
	if err != nil {
		return nil, err
	}
	allowed := false
	for _, role := range roles.List {
		if role.Status != enums.RoleStatusDisabled {
			allowed = true
			break
		}
	}
	if !allowed {
		return nil, bizerr.NewBizError(bizcode.CodeNoPermission, "无管理端访问权限")
	}

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
