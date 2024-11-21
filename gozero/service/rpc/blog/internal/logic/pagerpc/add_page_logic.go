package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/pagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPageLogic {
	return &AddPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建页面
func (l *AddPageLogic) AddPage(in *pagerpc.PageNewReq) (*pagerpc.PageDetails, error) {
	entity := convertPageIn(in)

	_, err := l.svcCtx.TPageModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertPageOut(entity), nil
}

func convertPageIn(in *pagerpc.PageNewReq) (out *model.TPage) {
	out = &model.TPage{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: jsonconv.AnyToJsonNE(in.CarouselCovers),
	}

	return out

}

func convertPageOut(in *model.TPage) (out *pagerpc.PageDetails) {
	out = &pagerpc.PageDetails{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: nil,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	jsonconv.JsonToAny(in.CarouselCovers, &out.CarouselCovers)
	return out
}
