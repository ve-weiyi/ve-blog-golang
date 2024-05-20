package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserMenusLogic {
	return &FindUserMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户菜单权限
func (l *FindUserMenusLogic) FindUserMenus(in *blog.UserReq) (*blog.MenuPageResp, error) {
	uid := in.UserId

	// 查用户
	//ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	//if err != nil {
	//	return nil, err
	//}

	// 查用户角色
	urs, err := l.svcCtx.UserRoleModel.FindALL(l.ctx, "user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range urs {
		ids = append(ids, v.RoleId)
	}

	// 查角色拥有的菜单
	rs, err := l.svcCtx.RoleMenuModel.FindALL(l.ctx, "role_id in (?)", ids)
	if err != nil {
		return nil, err
	}

	var mids []int64
	for _, v := range rs {
		mids = append(mids, v.MenuId)
	}

	// 查菜单信息
	list, err := l.svcCtx.MenuModel.FindALL(l.ctx, "id in (?)", mids)
	if err != nil {
		return nil, err
	}

	var root blog.MenuDetails
	root.Children = appendMenuChildren(&root, list)

	out := &blog.MenuPageResp{}
	out.Total = int64(len(root.Children))
	out.List = root.Children

	return out, nil
}

func appendMenuChildren(root *blog.MenuDetails, list []*model.Menu) (leafs []*blog.MenuDetails) {
	leafs = make([]*blog.MenuDetails, 0)
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertMenuModelToDetailPb(item)
			leaf.Children = appendMenuChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
