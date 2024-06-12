package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkLogic {
	return &FindFriendLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询友链
func (l *FindFriendLinkLogic) FindFriendLink(in *blog.IdReq) (*blog.FriendLink, error) {
	entity, err := l.svcCtx.FriendLinkModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkModelToPb(entity), nil
}
