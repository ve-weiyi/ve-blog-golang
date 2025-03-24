package friend

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建友链
func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.FriendNewReq) (resp *types.FriendBackVO, err error) {
	in := ConvertFriendPb(req)
	out, err := l.svcCtx.WebsiteRpc.AddFriend(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertFriendTypes(out)
	return resp, nil
}

func ConvertFriendPb(in *types.FriendNewReq) (out *websiterpc.FriendNewReq) {
	out = &websiterpc.FriendNewReq{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
	}

	return
}

func ConvertFriendTypes(in *websiterpc.FriendDetails) (out *types.FriendBackVO) {
	out = &types.FriendBackVO{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}

	return
}
