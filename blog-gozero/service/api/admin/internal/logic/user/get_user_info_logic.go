package user

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"

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

	return ConvertUserInfoTypes(info), nil
}

func ConvertUserInfoTypes(in *accountrpc.UserInfoResp) *types.UserInfoResp {
	roles := make([]*types.UserRoleLabel, 0)
	for _, v := range in.Roles {
		m := &types.UserRoleLabel{
			RoleId:    v.RoleId,
			RoleKey:   v.RoleKey,
			RoleLabel: v.RoleLabel,
		}

		roles = append(roles, m)
	}

	var info types.UserInfoExt
	jsonconv.JsonToAny(in.Info, &info)

	out := &types.UserInfoResp{
		UserId:      in.UserId,
		Username:    in.Username,
		Nickname:    in.Nickname,
		Avatar:      in.Avatar,
		Email:       in.Email,
		Phone:       in.Phone,
		Status:      in.Status,
		LoginType:   in.LoginType,
		IpAddress:   in.IpAddress,
		IpSource:    in.IpSource,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Roles:       roles,
		Perms:       make([]*types.UserApi, 0),
		UserInfoExt: info,
	}

	return out
}
