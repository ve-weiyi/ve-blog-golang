package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeleteRoleLogic) DeleteRole(in *account.IdReq) (*account.BatchResp, error) {
	result, err := l.svcCtx.RoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	result2, err := l.svcCtx.RoleModel.DeleteBatch(l.ctx, "parent_id = ? ", in.Id)
	if err != nil {
		return nil, err
	}

	return &account.BatchResp{
		SuccessCount: result + result2,
	}, nil
}
