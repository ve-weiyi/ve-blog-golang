package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChatRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteChatRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChatRecordListLogic {
	return &DeleteChatRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除聊天记录
func (l *DeleteChatRecordListLogic) DeleteChatRecordList(in *blog.IdsReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.ChatRecordModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
