package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type BatchDeleteTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteTagsLogic {
	return &BatchDeleteTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除标签
func (l *BatchDeleteTagsLogic) BatchDeleteTags(in *contentrpc.BatchDeleteTagsRequest) (*contentrpc.BatchDeleteTagsResponse, error) {
	var rows int64
	err := l.svcCtx.GormDB.Transaction(func(tx *gorm.DB) error {
		if _, err := l.svcCtx.TArticleTagModel.WithTx(tx).DeleteBatch(l.ctx, "tag_id in (?)", in.Ids); err != nil {
			return err
		}

		var err error
		rows, err = l.svcCtx.TTagModel.WithTx(tx).DeleteBatch(l.ctx, "id in (?)", in.Ids)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &contentrpc.BatchDeleteTagsResponse{SuccessCount: rows}, nil
}
