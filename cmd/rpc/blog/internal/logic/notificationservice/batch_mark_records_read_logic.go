package notificationservicelogic

import (
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/notificationrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type BatchMarkRecordsReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchMarkRecordsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchMarkRecordsReadLogic {
	return &BatchMarkRecordsReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchMarkRecordsReadLogic) BatchMarkRecordsRead(in *notificationrpc.BatchMarkRecordsReadRequest) (*notificationrpc.BatchMarkRecordsReadResponse, error) {
	if len(in.Ids) == 0 {
		return &notificationrpc.BatchMarkRecordsReadResponse{}, nil
	}

	uid, err := metax.GetUserIdFromCtx(l.ctx)
	if err != nil || uid == "" {
		return nil, bizerr.NewBizError(bizcode.CodeUnauthenticated, "缺少用户身份")
	}

	placeholders := make([]string, len(in.Ids))
	args := make([]any, 0, len(in.Ids)+3)
	args = append(args, uid, enums.NotifyChannelInbox, enums.NotifyRecordStatusUnread)
	for i, id := range in.Ids {
		placeholders[i] = "?"
		args = append(args, id)
	}

	fields := map[string]interface{}{
		"status":  enums.NotifyRecordStatusRead,
		"read_at": time.Now(),
	}
	rows, err := l.svcCtx.TNotifyRecordModel.UpdateFields(l.ctx, fields,
		"recipient = ? AND channel = ? AND status = ? AND id IN ("+strings.Join(placeholders, ",")+")", args...)
	if err != nil {
		return nil, err
	}

	return &notificationrpc.BatchMarkRecordsReadResponse{
		SuccessCount: rows,
	}, nil
}
