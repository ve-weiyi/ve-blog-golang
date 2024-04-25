package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWebsiteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新配置
func NewUpdateWebsiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWebsiteConfigLogic {
	return &UpdateWebsiteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWebsiteConfigLogic) UpdateWebsiteConfig(reqCtx *types.RestHeader, req *types.WebsiteConfig) (resp *types.EmptyResp, err error) {
	in := blog.SaveConfigReq{
		ConfigKey:   "website_config",
		ConfigValue: jsonconv.ObjectToJson(req),
	}

	_, err = l.svcCtx.ConfigRpc.SaveConfig(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return
}
