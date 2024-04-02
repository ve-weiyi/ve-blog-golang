package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebsiteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebsiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebsiteConfigLogic {
	return &GetWebsiteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebsiteConfigLogic) GetWebsiteConfig() (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
