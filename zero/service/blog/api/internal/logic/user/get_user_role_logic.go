package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleLogic {
	return &GetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoleLogic) GetUserRole(req *types.EmptyReq) (resp *types.UserRolesResp, err error) {
	in := convert.EmptyReq()

	out, err := l.svcCtx.UserRpc.GetUserRoles(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserRole
	jsonconv.ObjectMarshal(out.List, &list)

	resp = &types.UserRolesResp{}
	resp.List = list
	return
}
