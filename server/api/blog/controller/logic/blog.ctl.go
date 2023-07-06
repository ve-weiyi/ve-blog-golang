package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
	"github.com/ve-weiyi/ve-admin-store/server/infra/ws"
)

type BlogController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewBlogController(svcCtx *svc.ControllerContext) *BlogController {
	return &BlogController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// 创建api路由 https://www.jianshu.com/p/4bb4283632e4
//
//	@Tags			Blog
//	@Summary		标题
//	@Description	描述,可以有多个
//	@Security		ApiKeyAuth
//	@Param			file	formData	file								true	"上传文件"
//	@Param			id		path		int									true	"id"
//	@Param			token	header		string								true	"token"
//	@Param			data	body		entity.Api							true	"创建api路由"
//	@Success		200		{object}	response.Response{data=entity.Api}	"返回信息"
//	@Router			/api/version [get]
func (s *BlogController) ApiVersion(c *gin.Context) {

	s.ResponseOk(c, nil)
}

//	@Tags		Blog
//	@Summary	关于我
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	response.Response{data=entity.Api}	"返回信息"
//	@Router		/about [get]
func (s *BlogController) GetAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var api entity.Api
	err = s.ShouldBindJSON(c, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.CreateApi(reqCtx, &api)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Blog
//	@Summary	查询聊天记录
//	@Router		/ws [get]
func (s *BlogController) WebSocket(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 接收消息
	receive := func(msg []byte) {
		global.LOG.Println(string(msg))

		var chat entity.ChatRecord
		err = jsonconv.JsonToObject(string(msg), &chat)
		if err != nil {
			global.LOG.Error(err)
		}

		if chat.Content == "" {
			return
		}
		_, err = s.svcCtx.ChatRecordService.CreateChatRecord(reqCtx, &chat)
		if err != nil {
			global.LOG.Error(err)
		}
	}

	ws.HandleWebSocket(c.Writer, c.Request, receive)
}

//	@Tags		Blog
//	@Summary	查询聊天记录
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.PageInfo							true	"分页信息"
//	@Success	200		{object}	response.Response{data=entity.ChatRecord}	"返回信息"
//	@Router		/chat/records [post]
func (s *BlogController) FindChatRecords(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ChatRecordService.FindChatRecordList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}
