package photorpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAlbumLogic {
	return &AddAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建相册
func (l *AddAlbumLogic) AddAlbum(in *photorpc.AlbumNewReq) (*photorpc.AlbumDetails, error) {
	entity := convertAlbumIn(in)

	_, err := l.svcCtx.AlbumModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertAlbumOut(entity, nil), nil
}

func convertAlbumIn(in *photorpc.AlbumNewReq) (out *model.Album) {
	out = &model.Album{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
		CreatedAt:  time.Unix(in.CreatedAt, 0),
		UpdatedAt:  time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func convertAlbumOut(in *model.Album, cm map[int64]int) (out *photorpc.AlbumDetails) {
	var count int
	if v, ok := cm[in.Id]; ok {
		count = v
	}

	out = &photorpc.AlbumDetails{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt.Unix(),
		UpdatedAt:  in.UpdatedAt.Unix(),
		PhotoCount: int64(count),
	}

	return out
}
