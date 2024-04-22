package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
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
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.RoleRpc.DeleteRoleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResult{
		SuccessCount: out.SuccessCount,
	}, nil
}
