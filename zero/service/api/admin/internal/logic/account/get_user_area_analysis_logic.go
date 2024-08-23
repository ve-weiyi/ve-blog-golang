package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAreaAnalysisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户分布地区
func NewGetUserAreaAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAreaAnalysisLogic {
	return &GetUserAreaAnalysisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAreaAnalysisLogic) GetUserAreaAnalysis(req *types.EmptyReq) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
