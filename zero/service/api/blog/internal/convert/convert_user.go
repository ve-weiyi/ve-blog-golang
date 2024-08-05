package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertUserInfoTypes(in *blogrpc.UserInfoResp) (out *types.UserInfoResp) {
	out = &types.UserInfoResp{
		UserId:      in.UserId,
		Username:    in.Username,
		Nickname:    in.Nickname,
		Avatar:      in.Avatar,
		Email:       in.Email,
		Phone:       in.Phone,
		UserInfoExt: types.UserInfoExt{},
	}

	jsonconv.JsonToObject(in.Info, &out.UserInfoExt)

	return out
}
