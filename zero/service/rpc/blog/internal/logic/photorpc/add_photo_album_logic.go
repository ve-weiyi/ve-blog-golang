package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoAlbumLogic {
	return &AddPhotoAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建相册
func (l *AddPhotoAlbumLogic) AddPhotoAlbum(in *blog.PhotoAlbum) (*blog.PhotoAlbum, error) {
	entity := convert.ConvertPhotoAlbumPbToModel(in)

	_, err := l.svcCtx.PhotoAlbumModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumModelToPb(entity), nil
}
