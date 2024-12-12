package messagerpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatListLogic {
	return &FindChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询聊天记录列表
func (l *FindChatListLogic) FindChatList(in *messagerpc.FindChatListReq) (*messagerpc.FindChatListResp, error) {
	page, size, sorts, conditions, params := convertChatQuery(in)

	result, err := l.svcCtx.TChatModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.TChatModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*messagerpc.ChatDetails
	for _, v := range result {
		list = append(list, convertChatOut(v))
	}

	return &messagerpc.FindChatListResp{
		List:  list,
		Total: count,
	}, nil
}

func convertChatQuery(in *messagerpc.FindChatListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(1)
	size = int(in.Limit)

	if in.After != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "created_at >= ?"
		params = append(params, time.Unix(in.After, 0))
	}

	if in.Before != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "created_at <= ?"
		params = append(params, time.Unix(in.Before, 0))
	}

	if in.UserId != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "user_id = ?"
		params = append(params, in.UserId)
	}

	if in.Type != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "type = ?"
		params = append(params, in.Type)
	}

	if in.Status != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions = "status = ?"
		params = append(params, in.Status)
	}

	return
}
