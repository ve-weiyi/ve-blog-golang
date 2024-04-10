package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleResourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleResourceLogic {
	return &FindRoleResourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色
func (l *FindRoleResourceLogic) FindRoleResource(in *account.IdReq) (*account.RoleResourceResp, error) {
	// todo: add your logic here and delete this line

	return &account.RoleResourceResp{}, nil
}
