package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置密码
func (l *ResetPasswordLogic) ResetPassword(in *blog.ResetPasswordReq) (*blog.EmptyResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.ErrorInvalidParam
	}

	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	// 更新密码
	account.Password = crypto.BcryptHash(in.Password)

	_, err = l.svcCtx.UserAccountModel.Update(l.ctx, account)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
