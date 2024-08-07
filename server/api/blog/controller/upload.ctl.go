package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type UploadController struct {
	svcCtx *svc.ServiceContext
}

func NewUploadController(svcCtx *svc.ServiceContext) *UploadController {
	return &UploadController{
		svcCtx: svcCtx,
	}
}

// @Tags		Upload
// @Summary		上传文件
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		label	path		string									true	"标签"
// @Param		file	formData	file									true	"文件"
// @Success		200		{object}	response.Response{data=entity.UploadRecord}	"返回信息"
// @Router		/upload/{label} [post]
func (s *UploadController) UploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	// 获取上传的文件标签
	label := c.Param("label")

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewUploadService(s.svcCtx).UploadFile(reqCtx, label, file)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Upload
// @Summary		上传语言
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param 		type formData int true "消息类型"
// @Param 		file formData file true "语音文件"
// @Param 		user_id formData int true "用户ID"
// @Param 		nickname formData string true "用户昵称"
// @Param 		avatar formData string true "用户头像"
// @Param 		content formData string true "聊天内容"
// @Param 		created_at formData string true "创建时间"
// @Param 		ip_address formData string true "IP地址"
// @Param 		ip_source formData string true "IP来源"
// @Success		200		{object}	response.Response{data=entity.UploadRecord}	"返回信息"
// @Router		/voice [post]
func (s *UploadController) UploadVoice(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	tp := c.PostForm("type")
	//uid := c.PostForm("user_id")
	//nickname := c.PostForm("nickname")
	//avatar := c.PostForm("avatar")
	content := c.PostForm("content")
	//created_at := c.PostForm("created_at")
	//ip_address := c.PostForm("ip_address")
	//ip_source := c.PostForm("ip_source")
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	req := dto.VoiceVO{
		Type: cast.ToInt64(tp),
		//UserId:    cast.ToInt(uid),
		//Nickname:  nickname,
		//Avatar:    avatar,
		Content: content,
		//CreatedAt: time.Now(),
		//IPAddress: ip_address,
		//IPSource:  ip_source,
	}

	data, err := service.NewUploadService(s.svcCtx).UploadVoice(reqCtx, &req, file)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
