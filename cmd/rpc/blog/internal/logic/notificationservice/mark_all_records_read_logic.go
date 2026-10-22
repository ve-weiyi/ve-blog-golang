package notificationservicelogic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/notificationrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

type MarkAllRecordsReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkAllRecordsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkAllRecordsReadLogic {
	return &MarkAllRecordsReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkAllRecordsReadLogic) MarkAllRecordsRead(in *notificationrpc.MarkAllRecordsReadRequest) (*notificationrpc.MarkAllRecordsReadResponse, error) {
	fields := map[string]interface{}{
		"status":  enums.NotifyRecordStatusRead,
		"read_at": time.Now(),
	}

	rows, err := l.svcCtx.TNotifyRecordModel.UpdateFields(l.ctx, fields, "channel = ? AND recipient = ? AND status = ?", enums.NotifyChannelInbox, in.UserId, enums.NotifyRecordStatusUnread)
	if err != nil {
		return nil, err
	}

	return &notificationrpc.MarkAllRecordsReadResponse{
		SuccessCount: rows,
	}, nil
}
