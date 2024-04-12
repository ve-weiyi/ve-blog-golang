package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/rpcutils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户菜单权限
func (l *GetUserMenusLogic) GetUserMenus(in *account.EmptyReq) (*account.MenuPageResp, error) {
	uid, err := rpcutils.GetRPCInnerXUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 查用户
	ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	if err != nil {
		return nil, err
	}

	// 查用户角色
	urs, err := l.svcCtx.UserRoleModel.FindALL(l.ctx, "user_id = ?", ua.Id)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range urs {
		ids = append(ids, v.RoleId)
	}

	// 查角色拥有的接口
	rs, err := l.svcCtx.RoleMenuModel.FindALL(l.ctx, "id in (?)", ids)
	if err != nil {
		return nil, err
	}

	var apiIds []int64
	for _, v := range rs {
		apiIds = append(apiIds, v.MenuId)
	}

	// 查接口信息
	apis, err := l.svcCtx.MenuModel.FindALL(l.ctx, "id in (?)", apiIds)
	if err != nil {
		return nil, err
	}

	var list []*account.MenuDetailsDTO
	for _, v := range apis {
		list = append(list, convert.ConvertMenuModelToDetailPb(v))
	}

	out := &account.MenuPageResp{}
	out.Total = int64(len(list))
	out.List = list
	return out, nil
}
