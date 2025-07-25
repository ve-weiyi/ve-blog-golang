package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"

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

func (l *AddPageLogic) AddPage(req *types.PageNewReq) (resp *types.PageBackVO, err error) {
	in := ConvertPagePb(req)
	out, err := l.svcCtx.ResourceRpc.AddPage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertPageTypes(out)
	return resp, nil
}

func ConvertPagePb(in *types.PageNewReq) (out *resourcerpc.PageNewReq) {
	out = &resourcerpc.PageNewReq{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: in.CarouselCovers,
	}

	return
}

func ConvertPageTypes(in *resourcerpc.PageDetails) (out *types.PageBackVO) {
	out = &types.PageBackVO{
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
