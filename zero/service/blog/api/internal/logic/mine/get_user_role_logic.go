package mine

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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

func (l *GetUserRoleLogic) GetUserRole(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.UserRolesResp, err error) {
	in := &blog.UserReq{
		UserId: cast.ToInt64(reqCtx.HeaderXUserId),
	}

	out, err := l.svcCtx.UserRpc.GetUserRoles(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserRole
	jsonconv.ObjectToObject(out.List, &list)

	resp = &types.UserRolesResp{}
	resp.List = list
	return
}
