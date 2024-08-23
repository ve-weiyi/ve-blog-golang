package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindBannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBannerListLogic {
	return &FindBannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询页面列表
func (l *FindBannerListLogic) FindBannerList(in *photorpc.FindBannerListReq) (*photorpc.FindBannerListResp, error) {
	// todo: add your logic here and delete this line

	return &photorpc.FindBannerListResp{}, nil
}
