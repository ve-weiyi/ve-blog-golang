package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除页面
func NewDeletePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageLogic {
	return &DeletePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePageLogic) DeletePage(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
