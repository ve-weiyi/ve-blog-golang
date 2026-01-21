package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type NoticeController struct {
	svcCtx *svctx.ServiceContext
}

func NewNoticeController(svcCtx *svctx.ServiceContext) *NoticeController {
	return &NoticeController{
		svcCtx: svcCtx,
	}
}

// @Tags		Notice
// @Summary		"创建通知"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.AddNoticeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.NoticeBackVO}	"返回信息"
// @Router		/admin-api/v1/notice/add_notice [POST]
func (s *NoticeController) AddNotice(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.AddNoticeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).AddNotice(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"删除通知"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.BatchResp}	"返回信息"
// @Router		/admin-api/v1/notice/deletes_notice [DELETE]
func (s *NoticeController) DeletesNotice(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).DeletesNotice(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"分页获取通知列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.QueryNoticeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.PageResp}	"返回信息"
// @Router		/admin-api/v1/notice/find_notice_list [POST]
func (s *NoticeController) FindNoticeList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryNoticeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).FindNoticeList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"查询用户通知列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.QueryUserNoticeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.PageResp}	"返回信息"
// @Router		/admin-api/v1/notice/find_user_notice_list [POST]
func (s *NoticeController) FindUserNoticeList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryUserNoticeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).FindUserNoticeList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"查询通知详情"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.IdReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.NoticeBackVO}	"返回信息"
// @Router		/admin-api/v1/notice/get_notice [GET]
func (s *NoticeController) GetNotice(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).GetNotice(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"更新通知"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.UpdateNoticeReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.NoticeBackVO}	"返回信息"
// @Router		/admin-api/v1/notice/update_notice [PUT]
func (s *NoticeController) UpdateNotice(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateNoticeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).UpdateNotice(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Notice
// @Summary		"更新通知状态"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		types.UpdateNoticeStatusReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=types.NoticeBackVO}	"返回信息"
// @Router		/admin-api/v1/notice/update_notice_status [PUT]
func (s *NoticeController) UpdateNoticeStatus(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateNoticeStatusReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := logic.NewNoticeLogic(s.svcCtx).UpdateNoticeStatus(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
