package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoLogic {
	return &UpdatePhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePhotoLogic) UpdatePhoto(req *types.Photo) (resp *types.Photo, err error) {
	// todo: add your logic here and delete this line

	return
}
