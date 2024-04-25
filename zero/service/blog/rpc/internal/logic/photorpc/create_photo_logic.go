package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoLogic {
	return &CreatePhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建照片
func (l *CreatePhotoLogic) CreatePhoto(in *blog.Photo) (*blog.Photo, error) {
	// todo: add your logic here and delete this line

	return &blog.Photo{}, nil
}
