package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBannerLogic {
	return &DeleteBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除页面
func (l *DeleteBannerLogic) DeleteBanner(in *photorpc.IdReq) (*photorpc.BatchResp, error) {
	rows, err := l.svcCtx.BannerModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &photorpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
