package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex/metadata"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *blog.EmptyReq) (*blog.UserInfoResp, error) {
	uid, err := metadata.GetRPCInnerXUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	//ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	//if err != nil {
	//	return nil, err
	//}

	ui, err := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUserInfoModelToPb(ui), nil
}
