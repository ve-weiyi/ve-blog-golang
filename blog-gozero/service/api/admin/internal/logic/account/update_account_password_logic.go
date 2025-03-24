package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户密码
func NewUpdateAccountPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountPasswordLogic {
	return &UpdateAccountPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountPasswordLogic) UpdateAccountPassword(req *types.UpdateAccountPasswordReq) (resp *types.EmptyResp, err error) {
	in := &accountrpc.UpdateUserPasswordReq{
		UserId:   req.UserId,
		Password: req.Password,
	}

	_, err = l.svcCtx.AccountRpc.UpdateUserPassword(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
