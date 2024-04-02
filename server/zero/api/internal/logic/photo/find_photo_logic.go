package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoLogic {
	return &FindPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoLogic) FindPhoto(req *types.IdReq) (resp *types.Photo, err error) {
	// todo: add your logic here and delete this line

	return
}
