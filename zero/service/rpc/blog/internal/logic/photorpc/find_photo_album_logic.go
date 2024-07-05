package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumLogic {
	return &FindPhotoAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询相册
func (l *FindPhotoAlbumLogic) FindPhotoAlbum(in *blog.IdReq) (*blog.PhotoAlbum, error) {
	entity, err := l.svcCtx.PhotoAlbumModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumModelToPb(entity), nil
}
