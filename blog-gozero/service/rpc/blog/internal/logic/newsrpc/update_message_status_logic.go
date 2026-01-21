package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMessageStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageStatusLogic {
	return &UpdateMessageStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新留言状态
func (l *UpdateMessageStatusLogic) UpdateMessageStatus(in *newsrpc.UpdateMessageStatusReq) (*newsrpc.UpdateMessageStatusResp, error) {
	rows, err := l.svcCtx.TMessageModel.Updates(l.ctx, map[string]interface{}{
		"status": in.Status,
	}, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &newsrpc.UpdateMessageStatusResp{
		SuccessCount: rows,
	}, nil
}
