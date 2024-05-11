package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
func (l *FindRoleResourcesLogic) FindRoleResources(in *blog.IdReq) (*blog.RoleResourcesResp, error) {
	logx.Errorf("in: %v", in)
	ras, err := l.svcCtx.RoleApiModel.FindALL(l.ctx, "role_id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	rms, err := l.svcCtx.RoleMenuModel.FindALL(l.ctx, "role_id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	out := &blog.RoleResourcesResp{
		RoleId:  in.Id,
		ApiIds:  make([]int64, 0),
		MenuIds: make([]int64, 0),
	}

	for _, ra := range ras {
		out.ApiIds = append(out.ApiIds, ra.ApiId)
	}

	for _, rm := range rms {
		out.MenuIds = append(out.MenuIds, rm.MenuId)
	}

	return out, nil
}
