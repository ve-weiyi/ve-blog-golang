package notificationservicelogic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/notificationrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type MarkRecordReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkRecordReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkRecordReadLogic {
	return &MarkRecordReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 标记单条 inbox 已读（按投递记录 id）
func (l *MarkRecordReadLogic) MarkRecordRead(in *notificationrpc.MarkRecordReadRequest) (*notificationrpc.MarkRecordReadResponse, error) {
	fields := map[string]interface{}{
		"status":  enums.NotifyRecordStatusRead,
		"read_at": time.Now(),
	}

	uid, err := metax.GetUserIdFromCtx(l.ctx)
	if err != nil || uid == "" {
		return nil, bizerr.NewBizError(bizcode.CodeUnauthenticated, "缺少用户身份")
	}

	_, err = l.svcCtx.TNotifyRecordModel.UpdateFields(l.ctx, fields,
		"id = ? AND recipient = ? AND channel = ? AND status = ?",
		in.Id, uid, enums.NotifyChannelInbox, enums.NotifyRecordStatusUnread)
	if err != nil {
		return nil, err
	}

	return &notificationrpc.MarkRecordReadResponse{
		Success: true,
	}, nil
}
