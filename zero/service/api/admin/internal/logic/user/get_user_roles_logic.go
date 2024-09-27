package user

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户角色
func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.EmptyReq) (resp *types.UserRolesResp, err error) {
	in := &permissionrpc.UserIdReq{
		UserId: cast.ToInt64(l.ctx.Value("uid")),
	}

	out, err := l.svcCtx.PermissionRpc.FindUserRoles(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserRole
	jsonconv.ObjectToObject(out.List, &list)

	resp = &types.UserRolesResp{}
	resp.List = list
	return
}
