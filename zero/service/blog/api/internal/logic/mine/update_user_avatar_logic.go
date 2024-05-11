package mine

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAvatarLogic {
	return &UpdateUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAvatarLogic) UpdateUserAvatar(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
