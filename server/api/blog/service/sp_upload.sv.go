package service

import (
	"mime/multipart"
	"path"
	"time"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

type UploadService struct {
	svcCtx *svc.ServiceContext
}

func NewUploadService(svcCtx *svc.ServiceContext) *UploadService {
	return &UploadService{
		svcCtx: svcCtx,
	}
}

// 上传文件
func (l *UploadService) UploadFile(reqCtx *request.Context, label string, file *multipart.FileHeader) (data *entity.UploadRecord, err error) {
	glog.Println("上传文件")
	label = "upload" + label
	url, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.Uid), label), file)
	if err != nil {
		return nil, err
	}

	up := &entity.UploadRecord{
		UserId:   reqCtx.Uid,
		Label:    label,
		FileName: file.Filename,
		FileSize: file.Size,
		FileMd5:  crypto.Md5v(file.Filename, ""),
		FileUrl:  url,
	}

	return l.svcCtx.UploadRecordRepository.Create(reqCtx, up)
}

// 上传语言
func (l *UploadService) UploadVoice(reqCtx *request.Context, req *dto.VoiceVO, file *multipart.FileHeader) (data *entity.UploadRecord, err error) {
	label := "voice"
	filename := time.Now().Format("20060102150405") + ".mp3"

	glog.Println("上传语言")
	url, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.Uid), label), file)
	if err != nil {
		return nil, err
	}

	glog.Println("查询用户信息")
	user, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, reqCtx.Uid)
	if err != nil {
		return nil, err
	}

	glog.Println("创建聊天记录")
	var chat entity.ChatRecord
	chat.Type = constant.VoiceMessage
	chat.UserId = user.UserId
	chat.Avatar = user.Avatar
	chat.Nickname = user.Nickname
	chat.Content = url
	chat.CreatedAt = time.Now()
	chat.IpAddress = reqCtx.IpAddress
	chat.IpSource = reqCtx.GetIpSource()

	_, err = l.svcCtx.ChatRecordRepository.Create(reqCtx, &chat)
	if err != nil {
		return nil, err
	}

	glog.Println("Websocket广播")
	ws.Broadcast([]byte(jsonconv.ObjectToJson(chat)))

	glog.Println("创建上传记录")
	up := &entity.UploadRecord{
		UserId:   reqCtx.Uid,
		Label:    label,
		FileName: filename,
		FileSize: file.Size,
		FileMd5:  crypto.Md5v(filename, ""),
		FileUrl:  url,
	}
	return l.svcCtx.UploadRecordRepository.Create(reqCtx, up)
}
