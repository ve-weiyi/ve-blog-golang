package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"

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
	in := rolerpc.IdsReq{
		Ids: req.IDS,
	}

	_, err = l.svcCtx.RoleRpc.DeleteRoleList(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResult{}, nil
}
