package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

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
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.PhotoRpc.FindPhotoAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumTypes(out), nil
}
