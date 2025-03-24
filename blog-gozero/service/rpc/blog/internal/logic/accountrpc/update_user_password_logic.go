package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户密码
func (l *UpdateUserPasswordLogic) UpdateUserPassword(in *accountrpc.UpdateUserPasswordReq) (*accountrpc.EmptyResp, error) {
	// 验证用户是否存在
	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, bizerr.NewBizError(bizerr.CodeUserNotExist, err.Error())
	}

	// 更新密码
	user.Password = crypto.BcryptHash(in.Password)

	_, err = l.svcCtx.TUserModel.Save(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &accountrpc.EmptyResp{}, nil
}
