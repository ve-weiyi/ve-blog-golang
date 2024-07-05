package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询相册
func NewFindPhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumLogic {
	return &FindPhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumLogic) FindPhotoAlbum(req *types.IdReq) (resp *types.PhotoAlbum, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.PhotoRpc.FindPhotoAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumTypes(out), nil
}
