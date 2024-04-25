package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoLogic {
	return &FindPhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询照片
func (l *FindPhotoLogic) FindPhoto(in *blog.IdReq) (*blog.Photo, error) {
	// todo: add your logic here and delete this line

	return &blog.Photo{}, nil
}
