package logic

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

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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
	url, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.UID), label), file)
	if err != nil {
		return nil, err
	}

	up := &entity.UploadRecord{
		UserID:   reqCtx.UID,
		Label:    label,
		FileName: file.Filename,
		FileSize: int(file.Size),
		FileMd5:  crypto.MD5V([]byte(file.Filename)),
		FileURL:  url,
	}

	return l.svcCtx.UploadRecordRepository.Create(reqCtx, up)
}

// 上传语言
func (l *UploadService) UploadVoice(reqCtx *request.Context, req *request.VoiceVO, file *multipart.FileHeader) (data *entity.UploadRecord, err error) {
	label := "voice"
	filename := time.Now().Format("20060102150405") + ".mp3"

	glog.Println("上传语言")
	url, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.UID), label), file)
	if err != nil {
		return nil, err
	}

	glog.Println("查询用户信息")
	user, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, reqCtx.UID)
	if err != nil {
		return nil, err
	}

	glog.Println("创建聊天记录")
	var chat entity.ChatRecord
	chat.Type = constant.VoiceMessage
	chat.UserID = user.UserID
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
		UserID:   reqCtx.UID,
		Label:    label,
		FileName: filename,
		FileSize: int(file.Size),
		FileMd5:  crypto.MD5V([]byte(filename)),
		FileURL:  url,
	}
	return l.svcCtx.UploadRecordRepository.Create(reqCtx, up)
}
