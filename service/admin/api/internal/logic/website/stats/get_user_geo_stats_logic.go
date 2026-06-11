package stats

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/analyticsservice"
)

type GetUserGeoStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户地理分布
func NewGetUserGeoStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGeoStatsLogic {
	return &GetUserGeoStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserGeoStatsLogic) GetUserGeoStats(req *types.GetUserGeoStatsReq) (resp *types.GetUserGeoStatsResp, err error) {
	out, err := l.svcCtx.AnalyticsService.GetUserGeoStats(l.ctx, &analyticsservice.GetUserGeoStatsRequest{})
	if err != nil {
		return nil, err
	}

	resp = &types.GetUserGeoStatsResp{
		Users:    make([]*types.RegionStatVO, 0, len(out.Users)),
		Visitors: make([]*types.RegionStatVO, 0, len(out.Visitors)),
	}
	for _, r := range out.Users {
		resp.Users = append(resp.Users, &types.RegionStatVO{
			Name:  r.Name,
			Value: r.Value,
		})
	}
	for _, v := range out.Visitors {
		resp.Visitors = append(resp.Visitors, &types.RegionStatVO{
			Name:  v.Name,
			Value: v.Value,
		})
	}
	return
}
