package notificationservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/notificationrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type BatchDeleteNotifyRecordsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteNotifyRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteNotifyRecordsLogic {
	return &BatchDeleteNotifyRecordsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchDeleteNotifyRecordsLogic) BatchDeleteNotifyRecords(in *notificationrpc.BatchDeleteNotifyRecordsRequest) (*notificationrpc.BatchDeleteNotifyRecordsResponse, error) {
	if len(in.Ids) == 0 {
		return &notificationrpc.BatchDeleteNotifyRecordsResponse{}, nil
	}

	rows, err := l.svcCtx.TNotifyRecordModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &notificationrpc.BatchDeleteNotifyRecordsResponse{
		SuccessCount: rows,
	}, nil
}
