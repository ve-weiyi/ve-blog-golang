package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagLogic {
	return &FindTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagLogic) FindTag(req *types.IdReq) (resp *types.Tag, err error) {
	// todo: add your logic here and delete this line

	return
}
