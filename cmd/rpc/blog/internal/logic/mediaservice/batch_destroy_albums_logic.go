package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type BatchDestroyAlbumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDestroyAlbumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDestroyAlbumsLogic {
	return &BatchDestroyAlbumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除相册，连同相册下的照片一起删除。
func (l *BatchDestroyAlbumsLogic) BatchDestroyAlbums(in *mediarpc.BatchDestroyAlbumsRequest) (*mediarpc.BatchDestroyAlbumsResponse, error) {
	var rows int64
	err := l.svcCtx.GormDB.Transaction(func(tx *gorm.DB) error {
		if _, err := l.svcCtx.TPhotoModel.WithTx(tx).DeleteBatch(l.ctx, "album_id in (?)", in.Ids); err != nil {
			return err
		}

		var err error
		rows, err = l.svcCtx.TAlbumModel.WithTx(tx).DeleteBatch(l.ctx, "id in (?)", in.Ids)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &mediarpc.BatchDestroyAlbumsResponse{SuccessCount: rows}, nil
}
