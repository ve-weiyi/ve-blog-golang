package menurpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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
func (l *FindMenuListLogic) FindMenuList(in *blog.PageQuery) (*blog.FindMenuListResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.MenuModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var root blog.MenuDetails
	root.Children = appendMenuChildren(&root, result)

	out := &blog.FindMenuListResp{}
	out.List = root.Children

	return out, nil
}

func appendMenuChildren(root *blog.MenuDetails, list []*model.Menu) (leafs []*blog.MenuDetails) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertMenuModelToDetailPb(item)
			leaf.Children = appendMenuChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
