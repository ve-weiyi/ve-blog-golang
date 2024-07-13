package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertChatRecordPbToModel(in *blog.ChatRecord) (out *model.ChatRecord) {
	out = &model.ChatRecord{
		Id:        in.Id,
		UserId:    in.UserId,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Content:   in.Content,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		Type:      in.Type,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertChatRecordModelToPb(in *model.ChatRecord) (out *blog.ChatRecord) {
	out = &blog.ChatRecord{
		Id:        in.Id,
		UserId:    in.UserId,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Content:   in.Content,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		Type:      in.Type,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
