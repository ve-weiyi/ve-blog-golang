package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleApisLogic) UpdateRoleApis(reqCtx *types.RestHeader, req *types.UpdateRoleApisReq) (resp *types.EmptyResp, err error) {
	in := rolerpc.UpdateRoleApisReq{
		RoleId: req.RoleId,
		ApiIds: req.ApiIds,
	}

	_, err = l.svcCtx.RoleRpc.UpdateRoleApis(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
