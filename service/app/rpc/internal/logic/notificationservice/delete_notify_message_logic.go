package notificationservicelogic

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/notificationrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type DeleteNotifyMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteNotifyMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNotifyMessageLogic {
	return &DeleteNotifyMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除通知消息
func (l *DeleteNotifyMessageLogic) DeleteNotifyMessage(in *notificationrpc.DeleteNotifyMessageRequest) (*notificationrpc.DeleteNotifyMessageResponse, error) {
	if len(in.Ids) == 0 {
		return &notificationrpc.DeleteNotifyMessageResponse{}, nil
	}

	placeholders := make([]string, len(in.Ids))
	args := make([]interface{}, len(in.Ids))
	for i, id := range in.Ids {
		placeholders[i] = "?"
		args[i] = id
	}

	condition := fmt.Sprintf("id IN (%s)", strings.Join(placeholders, ","))
	rows, err := l.svcCtx.TNotifyMessageModel.DeleteBatch(l.ctx, condition, args...)
	if err != nil {
		return nil, err
	}

	return &notificationrpc.DeleteNotifyMessageResponse{
		SuccessCount: rows,
	}, nil
}
