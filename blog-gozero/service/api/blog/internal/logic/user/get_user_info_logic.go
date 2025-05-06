package user

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.EmptyReq) (resp *types.UserInfoResp, err error) {
	in := &accountrpc.UserIdReq{
		UserId: cast.ToString(l.ctx.Value(restx.HeaderUid)),
	}

	info, err := l.svcCtx.AccountRpc.GetUserInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	thp, err := l.svcCtx.AccountRpc.GetUserOauthInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return ConvertUserInfoTypes(info, thp), nil
}

func ConvertUserInfoTypes(in *accountrpc.UserInfoResp, thp *accountrpc.GetUserOauthInfoResp) (out *types.UserInfoResp) {
	var info types.UserInfoExt
	jsonconv.JsonToAny(in.Info, &info)

	thirdParty := make([]*types.UserThirdPartyInfo, 0)
	for _, v := range thp.List {
		thirdParty = append(thirdParty, &types.UserThirdPartyInfo{
			Platform:  v.Platform,
			OpenId:    v.OpenId,
			Nickname:  v.Nickname,
			Avatar:    v.Avatar,
			CreatedAt: v.CreatedAt,
		})
	}

	out = &types.UserInfoResp{
		UserId:       in.UserId,
		Username:     in.Username,
		Nickname:     in.Nickname,
		Avatar:       in.Avatar,
		Email:        in.Email,
		Phone:        in.Phone,
		RegisterType: in.RegisterType,
		CreatedAt:    in.CreatedAt,
		UserInfoExt:  info,
		ThirdParty:   thirdParty,
	}

	return out
}
