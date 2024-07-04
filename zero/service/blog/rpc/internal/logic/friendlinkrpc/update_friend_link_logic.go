package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLinkLogic {
	return &UpdateFriendLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新友链
func (l *UpdateFriendLinkLogic) UpdateFriendLink(in *blog.FriendLink) (*blog.FriendLink, error) {
	entity := convert.ConvertFriendLinkPbToModel(in)

	_, err := l.svcCtx.FriendLinkModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkModelToPb(entity), nil
}
