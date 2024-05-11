package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFriendLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFriendLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFriendLinkLogic {
	return &CreateFriendLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建友链
func (l *CreateFriendLinkLogic) CreateFriendLink(in *blog.FriendLink) (*blog.FriendLink, error) {
	entity := convert.ConvertFriendLinkPbToModel(in)

	result, err := l.svcCtx.FriendLinkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertFriendLinkModelToPb(result), nil
}
