package page

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"
)

type FindPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取页面列表
func NewFindPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageListLogic {
	return &FindPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageListLogic) FindPageList(req *types.PageQueryReq) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindPageListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	out, err := l.svcCtx.ResourceRpc.FindPageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.Page, 0)
	for _, v := range out.List {
		list = append(list, ConvertPageTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertPageTypes(in *resourcerpc.PageDetails) *types.Page {
	return &types.Page{
		Id:         in.Id,
		PageName:   in.PageName,
		PageLabel:  in.PageLabel,
		PageCover:  in.PageCover,
		IsCarousel: in.IsCarousel,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
}
