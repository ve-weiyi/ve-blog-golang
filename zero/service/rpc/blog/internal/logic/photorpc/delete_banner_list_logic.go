package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBannerListLogic {
	return &DeleteBannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除页面
func (l *DeleteBannerListLogic) DeleteBannerList(in *photorpc.IdsReq) (*photorpc.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &photorpc.BatchResp{}, nil
}
