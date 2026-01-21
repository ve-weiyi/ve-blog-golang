package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/blog/types"
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

// 分页获取留言列表
func (s *MessageLogic) FindMessageList(reqCtx *request.Context, in *types.QueryMessageReq) (out *types.PageResp, err error) {
	// todo

	return
}

// 创建留言
func (s *MessageLogic) AddMessage(reqCtx *request.Context, in *types.NewMessageReq) (out *types.EmptyResp, err error) {
	// todo

	return
}
