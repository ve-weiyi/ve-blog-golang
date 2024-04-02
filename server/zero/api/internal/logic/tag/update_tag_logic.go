package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.Tag) (resp *types.Tag, err error) {
	// todo: add your logic here and delete this line

	return
}
