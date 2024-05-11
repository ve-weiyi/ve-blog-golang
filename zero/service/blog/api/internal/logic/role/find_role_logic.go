package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
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

func (l *FindRoleLogic) FindRole(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.Role, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.RoleRpc.FindRole(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRoleTypes(out), nil
}
