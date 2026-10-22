package userservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/userrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type PatchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchUserLogic {
	return &PatchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户状态
func (l *PatchUserLogic) PatchUser(in *userrpc.PatchUserRequest) (*userrpc.PatchUserResponse, error) {
	if !enums.IsValidUserStatus(in.Status) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "用户状态不合法")
	}

	_, err := l.svcCtx.TUserModel.UpdateFields(l.ctx, map[string]interface{}{
		"status": in.Status,
	}, "user_id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	return &userrpc.PatchUserResponse{
		Success: true,
	}, nil
}
