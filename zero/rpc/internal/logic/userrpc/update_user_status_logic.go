package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户状态
func (l *UpdateUserStatusLogic) UpdateUserStatus(in *account.UpdateUserStatusReq) (*account.EmptyResp, error) {
	ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	ua.Status = in.Status

	_, err = l.svcCtx.UserAccountModel.Update(l.ctx, ua)
	if err != nil {
		return nil, err
	}

	return &account.EmptyResp{}, nil
}
