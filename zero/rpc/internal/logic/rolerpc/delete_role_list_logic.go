package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleListLogic {
	return &DeleteRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除角色
func (l *DeleteRoleListLogic) DeleteRoleList(in *account.IdsReq) (*account.BatchResult, error) {
	result, err := l.svcCtx.RoleModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &account.BatchResult{
		SuccessCount: result,
	}, nil
}
