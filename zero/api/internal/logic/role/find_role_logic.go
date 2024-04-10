package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"

	"github.com/zeromicro/go-zero/core/logx"
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
	in := rolerpc.IdReq{
		Id: req.ID,
	}

	out, err := l.svcCtx.RoleRpc.FindRole(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return convertRoleTypes(out), nil
}
