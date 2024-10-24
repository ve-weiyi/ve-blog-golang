package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindChatMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatMessageListLogic {
	return &FindChatMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询聊天记录列表
func (l *FindChatMessageListLogic) FindChatMessageList(in *chatrpc.FindChatMessageListReq) (*chatrpc.FindChatMessageListResp, error) {
	page, size, sorts, conditions, params := convertChatQuery(in)

	result, err := l.svcCtx.TChatMessageModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.TChatMessageModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*chatrpc.ChatMessageDetails
	for _, v := range result {
		list = append(list, convertChatMessageOut(v))
	}

	return &chatrpc.FindChatMessageListResp{
		List:  list,
		Total: count,
	}, nil
}

func convertChatQuery(in *chatrpc.FindChatMessageListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(1)
	size = int(in.Limit)
	sorts = "created_at desc"

	if in.After != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "created_at >= ?"
		params = append(params, in.After)
	}

	if in.Before != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "created_at <= ?"
		params = append(params, in.Before)
	}

	if in.ChatId != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "chat_id = ?"
		params = append(params, in.ChatId)
	}

	if in.UserId != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "user_id = ?"
		params = append(params, in.UserId)
	}

	if in.Type != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "type = ?"
		params = append(params, in.Type)
	}

	return
}
