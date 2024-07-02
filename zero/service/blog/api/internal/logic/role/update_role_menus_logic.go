package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenusLogic {
	return &UpdateRoleMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleMenusLogic) UpdateRoleMenus(req *types.UpdateRoleMenusReq) (resp *types.EmptyResp, err error) {
	in := rolerpc.UpdateRoleMenusReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}

	_, err = l.svcCtx.RoleRpc.UpdateRoleMenus(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
