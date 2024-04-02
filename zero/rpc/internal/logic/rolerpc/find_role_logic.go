package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleLogic {
	return &FindRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色
func (l *FindRoleLogic) FindRole(in *account.IdReq) (*account.RoleDetailsDTO, error) {
	result, err := l.svcCtx.RoleModel.First(l.ctx, "id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	return convertRoleModelToDetailPb(result), nil
}
