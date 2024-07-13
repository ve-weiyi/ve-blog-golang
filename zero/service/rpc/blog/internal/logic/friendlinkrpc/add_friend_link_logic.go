package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLinkLogic {
	return &AddFriendLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建友链
func (l *AddFriendLinkLogic) AddFriendLink(in *blog.FriendLink) (*blog.FriendLink, error) {
	entity := convert.ConvertFriendLinkPbToModel(in)

	_, err := l.svcCtx.FriendLinkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkModelToPb(entity), nil
}
