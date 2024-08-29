package friendrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/friendrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建友链
func (l *AddFriendLogic) AddFriend(in *friendrpc.FriendNew) (*friendrpc.FriendDetails, error) {
	entity := ConvertFriendIn(in)

	_, err := l.svcCtx.FriendModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertFriendOut(entity), nil
}

func ConvertFriendIn(in *friendrpc.FriendNew) (out *model.Friend) {
	out = &model.Friend{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		//CreatedAt:   time.Time{},
		//UpdatedAt:   time.Time{},
	}

	return out
}

func ConvertFriendOut(in *model.Friend) (out *friendrpc.FriendDetails) {
	out = &friendrpc.FriendDetails{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		CreatedAt:   in.CreatedAt.Unix(),
		UpdatedAt:   in.UpdatedAt.Unix(),
	}

	return out
}
