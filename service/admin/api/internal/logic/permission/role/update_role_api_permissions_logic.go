package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/permissionservice"
)

type UpdateRoleApiPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新角色接口权限
func NewUpdateRoleApiPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApiPermissionsLogic {
	return &UpdateRoleApiPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleApiPermissionsLogic) UpdateRoleApiPermissions(req *types.UpdateRoleApiPermissionsReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.PermissionService.UpdateRoleApi(l.ctx, &permissionservice.UpdateRoleApiRequest{
		RoleId: req.RoleId,
		ApiIds: req.ApiIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
