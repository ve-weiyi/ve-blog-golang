package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type BatchPatchPhotosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchPatchPhotosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchPatchPhotosLogic {
	return &BatchPatchPhotosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部分更新照片
func (l *BatchPatchPhotosLogic) BatchPatchPhotos(in *mediarpc.BatchPatchPhotosRequest) (*mediarpc.BatchPatchPhotosResponse, error) {
	if !enums.IsValidIsDelete(in.IsDelete) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "删除标记不合法")
	}

	rows, err := l.svcCtx.TPhotoModel.UpdateFields(l.ctx, map[string]interface{}{
		"is_delete": in.IsDelete,
	}, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &mediarpc.BatchPatchPhotosResponse{SuccessCount: rows}, nil
}
