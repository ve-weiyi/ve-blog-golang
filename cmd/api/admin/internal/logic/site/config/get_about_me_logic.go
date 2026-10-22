package config

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/siteservice"
	"github.com/ve-weiyi/blog-cloud/infra/domain/configkey"
)

type GetAboutMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取关于我的信息
func NewGetAboutMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAboutMeLogic {
	return &GetAboutMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAboutMeLogic) GetAboutMe(req *types.EmptyReq) (resp *types.AboutMeVO, err error) {
	out, err := l.svcCtx.SiteService.GetConfig(l.ctx, &siteservice.GetConfigRequest{
		ConfigKey: configkey.AboutMe,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.AboutMeVO{
		Content: out.ConfigValue,
	}
	return
}
