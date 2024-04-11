package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindApiListLogic) FindApiList(req *types.PageQuery) (resp *types.PageResult, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.ApiRpc.FindApiList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ApiDetailsDTO
	for _, role := range out.List {
		list = append(list, convert.ConvertApiDetailsTypes(role))
	}

	resp = &types.PageResult{}
	resp.Page = in.Limit.Page
	resp.PageSize = in.Limit.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
