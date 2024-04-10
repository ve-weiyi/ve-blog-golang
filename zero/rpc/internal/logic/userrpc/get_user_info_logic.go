package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/rpcutils"

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
func (l *GetUserInfoLogic) GetUserInfo(in *account.EmptyReq) (*account.UserInfoResp, error) {
	uid, err := rpcutils.GetRPCInnerXUserId(l.ctx)
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
