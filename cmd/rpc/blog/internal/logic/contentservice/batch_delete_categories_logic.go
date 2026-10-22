package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type BatchDeleteCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteCategoriesLogic {
	return &BatchDeleteCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除分类。分类下的文章保留，统一归入未分类（category_id = 0）。
func (l *BatchDeleteCategoriesLogic) BatchDeleteCategories(in *contentrpc.BatchDeleteCategoriesRequest) (*contentrpc.BatchDeleteCategoriesResponse, error) {
	var rows int64
	err := l.svcCtx.GormDB.Transaction(func(tx *gorm.DB) error {
		if _, err := l.svcCtx.TArticleModel.WithTx(tx).UpdateFields(l.ctx,
			map[string]interface{}{"category_id": 0},
			"category_id in (?)", in.Ids); err != nil {
			return err
		}

		var err error
		rows, err = l.svcCtx.TCategoryModel.WithTx(tx).DeleteBatch(l.ctx, "id in (?)", in.Ids)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &contentrpc.BatchDeleteCategoriesResponse{SuccessCount: rows}, nil
}
