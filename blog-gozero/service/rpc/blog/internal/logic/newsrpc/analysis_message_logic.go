package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisMessageLogic {
	return &AnalysisMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 消息数据分析
func (l *AnalysisMessageLogic) AnalysisMessage(in *newsrpc.AnalysisMessageReq) (*newsrpc.AnalysisMessageResp, error) {
	rc, err := l.svcCtx.TMessageModel.FindCount(l.ctx, "")
	if err != nil {
		return nil, err
	}

	return &newsrpc.AnalysisMessageResp{
		MessageCount: rc,
	}, nil
}
