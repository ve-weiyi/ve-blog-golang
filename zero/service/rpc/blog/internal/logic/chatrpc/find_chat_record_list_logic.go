package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindChatRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatRecordListLogic {
	return &FindChatRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询聊天记录列表
func (l *FindChatRecordListLogic) FindChatRecordList(in *chatrpc.FindChatRecordListReq) (*chatrpc.FindChatRecordListResp, error) {
	// todo: add your logic here and delete this line

	return &chatrpc.FindChatRecordListResp{}, nil
}
