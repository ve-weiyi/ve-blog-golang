package authrpclogic

import (
	"context"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoffLogic {
	return &LogoffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销
func (l *LogoffLogic) Logoff(in *blog.LogoffReq) (*blog.EmptyResp, error) {
	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	err = l.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		_, err = l.logoff(l.ctx, tx, account.Id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}

func (l *LogoffLogic) logoff(ctx context.Context, tx *gorm.DB, uid int64) (*blog.EmptyResp, error) {
	// 删除用户账号
	_, err := l.svcCtx.UserAccountModel.WithTransaction(tx).Delete(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 删除用户信息
	_, err = l.svcCtx.UserInformationModel.WithTransaction(tx).DeleteBatch(ctx, "user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	// 删除用户角色
	_, err = l.svcCtx.UserRoleModel.WithTransaction(tx).DeleteBatch(ctx, "user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
