package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取留言列表
func NewGetRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarkListLogic {
	return &GetRemarkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRemarkListLogic) GetRemarkList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.RemarkRpc.FindRemarkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.RemarkRpc.FindRemarkCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Remark
	for _, v := range out.List {
		list = append(list, convert.ConvertRemarkTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
