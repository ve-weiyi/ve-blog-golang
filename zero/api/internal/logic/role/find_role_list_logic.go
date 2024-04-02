package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
)

type FindRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListLogic {
	return &FindRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleListLogic) FindRoleList(req *types.PageQuery) (resp []types.Role, err error) {
	//in:= rolerpc.PageQuery{
	//	Limit:  req.Limit,
	//	Sorts: nil,
	//	Conditions: make([]*account.PageCondition, 0),
	//}

	return
}
