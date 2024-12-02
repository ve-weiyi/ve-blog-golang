package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
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
	in := &websiterpc.FindPageListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	out, err := l.svcCtx.WebsiteRpc.FindPageList(l.ctx, in)
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

func ConvertPageTypes(in *websiterpc.PageDetails) *types.Page {
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
