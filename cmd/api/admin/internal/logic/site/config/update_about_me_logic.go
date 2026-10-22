package config

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/siteservice"
	"github.com/ve-weiyi/blog-cloud/infra/domain/configkey"
)

type UpdateAboutMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新关于我的信息
func NewUpdateAboutMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAboutMeLogic {
	return &UpdateAboutMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAboutMeLogic) UpdateAboutMe(req *types.AboutMeVO) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.SiteService.SaveConfig(l.ctx, &siteservice.SaveConfigRequest{
		ConfigKey:   configkey.AboutMe,
		ConfigValue: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
