package mine

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(reqCtx *types.RestHeader, req *types.UserInfoReq) (resp *types.EmptyResp, err error) {
	in := &blog.UpdateUserInfoReq{
		UserId:   cast.ToInt64(reqCtx.HeaderXUserId),
		Nickname: req.Nickname,
		//Phone:    req.Phone,
		Website: req.Website,
		Intro:   req.Intro,
	}

	_, err = l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
