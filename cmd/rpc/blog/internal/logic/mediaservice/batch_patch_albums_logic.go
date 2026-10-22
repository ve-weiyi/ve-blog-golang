package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type BatchPatchAlbumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchPatchAlbumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchPatchAlbumsLogic {
	return &BatchPatchAlbumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部分更新相册。删除相册时级联删除照片；恢复相册不反向恢复照片，
// 避免把此前单独删除的照片一并复活。
func (l *BatchPatchAlbumsLogic) BatchPatchAlbums(in *mediarpc.BatchPatchAlbumsRequest) (*mediarpc.BatchPatchAlbumsResponse, error) {
	if !enums.IsValidIsDelete(in.IsDelete) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "删除标记不合法")
	}

	var rows int64
	err := l.svcCtx.GormDB.Transaction(func(tx *gorm.DB) error {
		var err error
		rows, err = l.svcCtx.TAlbumModel.WithTx(tx).UpdateFields(l.ctx, map[string]interface{}{
			"is_delete": in.IsDelete,
		}, "id in (?)", in.Ids)
		if err != nil {
			return err
		}

		if in.IsDelete == 1 {
			_, err = l.svcCtx.TPhotoModel.WithTx(tx).UpdateFields(l.ctx, map[string]interface{}{
				"is_delete": in.IsDelete,
			}, "album_id in (?)", in.Ids)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &mediarpc.BatchPatchAlbumsResponse{SuccessCount: rows}, nil
}
