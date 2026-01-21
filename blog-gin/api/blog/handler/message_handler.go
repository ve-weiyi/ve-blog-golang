package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/blog/logic"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/blog/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type MessageController struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageController(svcCtx *svctx.ServiceContext) *MessageController {
	return &MessageController{
		svcCtx: svcCtx,
	}
}

// @Tags		Message
// @Summary		"分页获取留言列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.QueryMessageReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.PageResp}	"返回信息"
// @Router		/blog-api/v1/message/find_message_list [POST]
func (s *MessageController) FindMessageList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewMessageLogic(s.svcCtx).FindMessageList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Message
// @Summary		"创建留言"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.NewMessageReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/message/add_message [POST]
func (s *MessageController) AddMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.NewMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewMessageLogic(s.svcCtx).AddMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
