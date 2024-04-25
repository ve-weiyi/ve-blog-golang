package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertRemarkModelToPb(in *model.Remark) (out *blog.Remark) {
	out = &blog.Remark{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}

func ConvertRemarkPbToModel(in *blog.Remark) (out *model.Remark) {
	out = &model.Remark{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      time.Unix(in.CreatedAt, 0),
		UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}
