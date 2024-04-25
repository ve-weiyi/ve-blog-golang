package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建照片
func NewCreatePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoLogic {
	return &CreatePhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePhotoLogic) CreatePhoto(reqCtx *types.RestHeader, req *types.Photo) (resp *types.Photo, err error) {
	// todo: add your logic here and delete this line

	return
}
