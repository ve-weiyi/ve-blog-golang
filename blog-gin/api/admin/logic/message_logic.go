package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type MessageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageLogic(svcCtx *svctx.ServiceContext) *MessageLogic {
	return &MessageLogic{
		svcCtx: svcCtx,
	}
}

// 删除留言
func (s *MessageLogic) DeletesMessage(reqCtx *request.Context, in *types.IdsReq) (out *types.BatchResp, err error) {
	// todo

	return
}

// 分页获取留言列表
func (s *MessageLogic) FindMessageList(reqCtx *request.Context, in *types.QueryMessageReq) (out *types.PageResp, err error) {
	// todo

	return
}

// 更新留言状态
func (s *MessageLogic) UpdateMessageStatus(reqCtx *request.Context, in *types.UpdateMessageStatusReq) (out *types.BatchResp, err error) {
	// todo

	return
}
