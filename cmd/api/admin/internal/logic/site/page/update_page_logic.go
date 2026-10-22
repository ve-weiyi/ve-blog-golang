package page

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/siteservice"
)

type UpdatePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新页面
func NewUpdatePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePageLogic {
	return &UpdatePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePageLogic) UpdatePage(req *types.UpdatePageReq) (resp *types.PageVO, err error) {
	out, err := l.svcCtx.SiteService.UpdatePage(l.ctx, &siteservice.UpdatePageRequest{
		Id:             req.Id,
		PageName:       req.PageName,
		PageLabel:      req.PageLabel,
		PageCover:      req.PageCover,
		IsCarousel:     req.IsCarousel,
		CarouselCovers: req.CarouselCovers,
	})
	if err != nil {
		return nil, err
	}

	return &types.PageVO{
		Id:             out.Page.Id,
		PageName:       out.Page.PageName,
		PageLabel:      out.Page.PageLabel,
		PageCover:      out.Page.PageCover,
		IsCarousel:     out.Page.IsCarousel,
		CarouselCovers: out.Page.CarouselCovers,
		CreatedAt:      out.Page.CreatedAt,
		UpdatedAt:      out.Page.UpdatedAt,
	}, nil
}
