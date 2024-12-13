package website

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 访客上报
func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportLogic) Report(req *types.EmptyReq) (resp *types.ReportResp, err error) {
	terminal := cast.ToString(l.ctx.Value("terminal"))

	if terminal == "" {
		tourist, err := l.svcCtx.WebsiteRpc.GetTouristInfo(l.ctx, &websiterpc.EmptyReq{})
		if err != nil {
			return nil, err
		}

		terminal = tourist.TouristId
	}

	_, err = l.svcCtx.WebsiteRpc.AddVisit(l.ctx, &websiterpc.AddVisitReq{Visitor: terminal})
	if err != nil {
		return nil, err
	}

	return &types.ReportResp{
		TerminalId: terminal,
	}, nil
}
