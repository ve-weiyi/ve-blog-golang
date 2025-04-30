package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitTrendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取访客数据趋势
func NewGetVisitTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitTrendLogic {
	return &GetVisitTrendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitTrendLogic) GetVisitTrend(req *types.GetVisitTrendReq) (resp *types.GetVisitTrendResp, err error) {
	// 查询用户浏览量
	visit, err := l.svcCtx.WebsiteRpc.FindVisitTrend(l.ctx, &websiterpc.FindVisitTrendReq{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, err
	}

	uvs := make([]*types.VisitTrendVO, 0)
	for _, v := range visit.UvTrend {
		m := &types.VisitTrendVO{
			Date:  v.Date,
			Count: v.Count,
		}

		uvs = append(uvs, m)
	}

	pvs := make([]*types.VisitTrendVO, 0)
	for _, v := range visit.PvTrend {
		m := &types.VisitTrendVO{
			Date:  v.Date,
			Count: v.Count,
		}

		pvs = append(pvs, m)
	}

	return &types.GetVisitTrendResp{
		UvTrend: uvs,
		PvTrend: pvs,
	}, nil
}
