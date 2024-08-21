package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除相册
func NewDeletePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoAlbumLogic {
	return &DeletePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoAlbumLogic) DeletePhotoAlbum(req *types.IdReq) (resp *types.BatchResp, err error) {
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.PhotoRpc.DeletePhotoAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
