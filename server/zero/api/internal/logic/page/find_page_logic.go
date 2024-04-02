package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageLogic {
	return &FindPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageLogic) FindPage(req *types.IdReq) (resp *types.Page, err error) {
	// todo: add your logic here and delete this line

	return
}
