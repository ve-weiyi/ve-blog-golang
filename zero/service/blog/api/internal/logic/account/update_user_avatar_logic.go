<<<<<<<< HEAD:zero/service/api/blog/internal/logic/user/update_user_avatar_logic.go
package user
========
package account
>>>>>>>> 09fef341 (v2.0.0 修改项目结构 (#17)):zero/service/blog/api/internal/logic/account/update_user_avatar_logic.go

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户信息
func NewUpdateUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAvatarLogic {
	return &UpdateUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

<<<<<<<< HEAD:zero/service/api/blog/internal/logic/user/update_user_avatar_logic.go
func (l *UpdateUserAvatarLogic) UpdateUserAvatar(req *types.UpdateUserAvatarReq) (resp *types.EmptyResp, err error) {
========
func (l *UpdateUserAvatarLogic) UpdateUserAvatar(req *types.EmptyReq) (resp *types.EmptyResp, err error) {
>>>>>>>> 09fef341 (v2.0.0 修改项目结构 (#17)):zero/service/blog/api/internal/logic/account/update_user_avatar_logic.go
	// todo: add your logic here and delete this line

	return
}
