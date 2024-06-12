package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoAlbumLogic {
	return &UpdatePhotoAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新相册
func (l *UpdatePhotoAlbumLogic) UpdatePhotoAlbum(in *blog.PhotoAlbum) (*blog.PhotoAlbum, error) {
	entity := convert.ConvertPhotoAlbumPbToModel(in)

	_, err := l.svcCtx.PhotoAlbumModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumModelToPb(entity), nil
}
