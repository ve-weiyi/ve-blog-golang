package album

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/mediaservice"
)

type UpdateAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新相册
func NewUpdateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlbumLogic {
	return &UpdateAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAlbumLogic) UpdateAlbum(req *types.UpdateAlbumReq) (resp *types.AlbumVO, err error) {
	out, err := l.svcCtx.MediaService.UpdateAlbum(l.ctx, &mediaservice.UpdateAlbumRequest{
		Id:         req.Id,
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.AlbumVO{
		Id:         out.Album.Id,
		AlbumName:  out.Album.AlbumName,
		AlbumDesc:  out.Album.AlbumDesc,
		AlbumCover: out.Album.AlbumCover,
		IsDelete:   out.Album.IsDelete,
		Status:     out.Album.Status,
		PhotoCount: out.Album.PhotoCount,
		CreatedAt:  out.Album.CreatedAt,
		UpdatedAt:  out.Album.UpdatedAt,
	}, nil
}
