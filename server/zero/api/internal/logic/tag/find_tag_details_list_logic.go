package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindTagDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagDetailsListLogic {
	return &FindTagDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagDetailsListLogic) FindTagDetailsList(req *types.PageQuery) (resp []types.TagDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
