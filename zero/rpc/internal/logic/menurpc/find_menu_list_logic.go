package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取菜单列表
func (l *FindMenuListLogic) FindMenuList(in *account.PageQuery) (*account.MenuPageResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.MenuModel.FindList(l.ctx, page, size, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.MenuModel.Count(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	var root account.MenuDetailsDTO
	root.Children = appendMenuChildren(&root, result)

	out := &account.MenuPageResp{}
	out.Total = total
	out.List = root.Children

	return out, nil
}

func appendMenuChildren(root *account.MenuDetailsDTO, list []*model.Menu) (leafs []*account.MenuDetailsDTO) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertMenuModelToDetailPb(item)
			leaf.Children = appendMenuChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
