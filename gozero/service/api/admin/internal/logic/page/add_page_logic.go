package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/pagerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建页面
func NewAddPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPageLogic {
	return &AddPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPageLogic) AddPage(req *types.PageNewReq) (resp *types.PageBackDTO, err error) {
	in := ConvertPagePb(req)
	out, err := l.svcCtx.PageRpc.AddPage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertPageTypes(out)
	return resp, nil
}

func ConvertPagePb(in *types.PageNewReq) (out *pagerpc.PageNewReq) {
	out = &pagerpc.PageNewReq{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: in.CarouselCovers,
	}

	return
}

func ConvertPageTypes(in *pagerpc.PageDetails) (out *types.PageBackDTO) {
	out = &types.PageBackDTO{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: in.CarouselCovers,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}

	return
}
