package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleListLogic {
	return &DeleteRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleListLogic) DeleteRoleList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
