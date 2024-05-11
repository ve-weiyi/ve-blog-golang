package mine

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.UserInfoResp, err error) {
	in := &blog.UserReq{
		UserId: cast.ToInt64(reqCtx.HeaderXUserId),
	}

	info, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResp{
		Id:        info.Id,
		UserId:    info.UserId,
		Email:     info.Email,
		Nickname:  info.Nickname,
		Avatar:    info.Avatar,
		Phone:     info.Phone,
		Intro:     info.Intro,
		Website:   info.Website,
		CreatedAt: 0,
		UpdatedAt: 0,
	}

	return
}
