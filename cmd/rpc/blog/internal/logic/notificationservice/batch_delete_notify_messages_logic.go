package notificationservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/notificationrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type BatchDeleteNotifyMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteNotifyMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteNotifyMessagesLogic {
	return &BatchDeleteNotifyMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除通知消息，连同对应的投递记录一起删除。
func (l *BatchDeleteNotifyMessagesLogic) BatchDeleteNotifyMessages(in *notificationrpc.BatchDeleteNotifyMessagesRequest) (*notificationrpc.BatchDeleteNotifyMessagesResponse, error) {
	if len(in.Ids) == 0 {
		return &notificationrpc.BatchDeleteNotifyMessagesResponse{}, nil
	}

	var rows int64
	err := l.svcCtx.GormDB.Transaction(func(tx *gorm.DB) error {
		if _, err := l.svcCtx.TNotifyRecordModel.WithTx(tx).DeleteBatch(l.ctx,
			"message_id in (?)", in.Ids); err != nil {
			return err
		}

		var err error
		rows, err = l.svcCtx.TNotifyMessageModel.WithTx(tx).DeleteBatch(l.ctx,
			"id in (?)", in.Ids)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &notificationrpc.BatchDeleteNotifyMessagesResponse{
		SuccessCount: rows,
	}, nil
}
