package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleResourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleResourcesLogic {
	return &FindRoleResourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色
func (l *FindRoleResourcesLogic) FindRoleResources(in *account.IdReq) (*account.RoleResourcesResp, error) {
	ras, err := l.svcCtx.RoleApiModel.FindALL(l.ctx, "role_id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	rms, err := l.svcCtx.RoleMenuModel.FindALL(l.ctx, "role_id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	out := &account.RoleResourcesResp{}

	for _, ra := range ras {
		out.ApiIdList = append(out.ApiIdList, ra.ApiId)
	}

	for _, rm := range rms {
		out.MenuIdList = append(out.MenuIdList, rm.MenuId)
	}

	return out, nil
}
