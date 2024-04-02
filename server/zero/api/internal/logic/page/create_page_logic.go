package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePageLogic {
	return &CreatePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePageLogic) CreatePage(req *types.Page) (resp *types.Page, err error) {
	// todo: add your logic here and delete this line

	return
}
