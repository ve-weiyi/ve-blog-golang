package apirpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

type FindApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取接口列表
func (l *FindApiListLogic) FindApiList(in *blog.PageQuery) (*blog.ApiPageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.ApiModel.FindList(l.ctx, limit, offset, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var root blog.ApiDetails
	root.Children = appendApiChildren(&root, result)

	out := &blog.ApiPageResp{}
	out.Total = int64(len(root.Children))
	out.List = root.Children

	return out, nil
}

func appendApiChildren(root *blog.ApiDetails, list []*model.Api) (leafs []*blog.ApiDetails) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertApiModelToDetailPb(item)
			leaf.Children = appendApiChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
