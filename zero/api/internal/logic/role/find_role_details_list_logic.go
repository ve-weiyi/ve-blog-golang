package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleDetailsListLogic {
	return &FindRoleDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleDetailsListLogic) FindRoleDetailsList(req *types.PageQuery) (resp []types.RoleDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
