package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisUserLogic {
	return &AnalysisUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分布区域
func (l *AnalysisUserLogic) AnalysisUser(in *accountrpc.EmptyReq) (*accountrpc.AnalysisUserResp, error) {
	uc, err := l.svcCtx.TUserModel.FindCount(l.ctx, "")
	if err != nil {
		return nil, err
	}

	return &accountrpc.AnalysisUserResp{
		UserCount: uc,
	}, nil
}
