package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertUserInfoTypes(in *blog.UserInfoResp) (out *types.UserInfoResp) {
	out = &types.UserInfoResp{
		UserId:      in.UserId,
		Username:    in.Username,
		Nickname:    in.Nickname,
		Avatar:      in.Avatar,
		UserInfoExt: types.UserInfoExt{},
	}

	jsonconv.JsonToObject(in.Info, &out.UserInfoExt)

	return out
}
