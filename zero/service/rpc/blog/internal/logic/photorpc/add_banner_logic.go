package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBannerLogic {
	return &AddBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建页面
func (l *AddBannerLogic) AddBanner(in *photorpc.BannerNew) (*photorpc.BannerDetails, error) {
	// todo: add your logic here and delete this line

	return &photorpc.BannerDetails{}, nil
}
