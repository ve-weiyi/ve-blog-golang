package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoAlbumLogic {
	return &CreatePhotoAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建相册
func (l *CreatePhotoAlbumLogic) CreatePhotoAlbum(in *blog.PhotoAlbum) (*blog.PhotoAlbum, error) {
	entity := convert.ConvertPhotoAlbumPbToModel(in)

	result, err := l.svcCtx.PhotoAlbumModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumModelToPb(result), nil
}
