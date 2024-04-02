package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageLogic {
	return &DeletePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePageLogic) DeletePage(req *types.IdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
