package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

type FindRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleLogic {
	return &FindRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleLogic) FindRole(req *types.IdReq) (resp *types.Role, err error) {
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.RoleRpc.FindRole(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRoleTypes(out), nil
}
