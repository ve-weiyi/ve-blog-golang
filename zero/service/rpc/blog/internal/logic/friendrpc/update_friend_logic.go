package friendrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/friendrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLogic {
	return &UpdateFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新友链
func (l *UpdateFriendLogic) UpdateFriend(in *friendrpc.FriendNew) (*friendrpc.FriendDetails, error) {
	entity := ConvertFriendIn(in)

	_, err := l.svcCtx.FriendModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertFriendOut(entity), nil
}
