package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
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

func (l *DeleteRoleListLogic) DeleteRoleList(reqCtx *types.RestHeader, req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.RoleRpc.DeleteRoleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
