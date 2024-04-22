package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *FindApiListLogic) FindApiList(in *account.PageQuery) (*account.ApiPageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.ApiModel.FindList(l.ctx, limit, offset, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	var root account.ApiDetailsDTO
	root.Children = appendApiChildren(&root, result)

	out := &account.ApiPageResp{}
	out.Total = int64(len(root.Children))
	out.List = root.Children

	return out, nil
}

func appendApiChildren(root *account.ApiDetailsDTO, list []*model.Api) (leafs []*account.ApiDetailsDTO) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := convert.ConvertApiModelToDetailPb(item)
			leaf.Children = appendApiChildren(leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
