package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertOperationLogPb(in *types.OperationLog) (out *blogrpc.OperationLog) {
	return &blogrpc.OperationLog{
		Id:             in.Id,
		UserId:         in.UserId,
		Nickname:       in.Nickname,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptModule,
		OptDesc:        in.OptDesc,
		RequestUrl:     in.RequestUrl,
		RequestMethod:  in.RequestMethod,
		RequestHeader:  in.RequestHeader,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}
}

func ConvertOperationLogTypes(in *blogrpc.OperationLog) (out *types.OperationLog) {

	return &types.OperationLog{
		Id:             in.Id,
		UserId:         in.UserId,
		Nickname:       in.Nickname,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptModule,
		OptDesc:        in.OptDesc,
		RequestUrl:     in.RequestUrl,
		RequestMethod:  in.RequestMethod,
		RequestHeader:  in.RequestHeader,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}
}
