package convert

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertOperationLogPbToModel(in *blog.OperationLog) (out *model.OperationLog) {
	out = &model.OperationLog{
		Id:             in.Id,
		UserId:         in.UserId,
		Nickname:       in.Nickname,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptDesc,
		OptDesc:        in.OptDesc,
		RequestUrl:     in.RequestUrl,
		RequestMethod:  in.RequestMethod,
		RequestHeader:  in.RequestHeader,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		CreatedAt:      time.Unix(in.CreatedAt, 0),
		UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertOperationLogModelToPb(in *model.OperationLog) (out *blog.OperationLog) {
	out = &blog.OperationLog{
		Id:             in.Id,
		UserId:         in.UserId,
		Nickname:       in.Nickname,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptDesc,
		OptDesc:        in.OptDesc,
		RequestUrl:     in.RequestUrl,
		RequestMethod:  in.RequestMethod,
		RequestHeader:  in.RequestHeader,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}
