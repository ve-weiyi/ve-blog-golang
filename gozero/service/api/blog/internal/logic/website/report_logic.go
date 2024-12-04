package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ReportLogic) Report(req *types.ReportReq) (resp *types.ReportResp, err error) {
	in := &websiterpc.EmptyReq{}

	out, err := l.svcCtx.WebsiteRpc.Report(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.ReportResp{
		TerminalId: out.Visitor,
	}, nil
}
