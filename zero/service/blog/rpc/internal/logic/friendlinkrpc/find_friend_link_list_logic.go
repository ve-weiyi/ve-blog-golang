package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendLinkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFriendLinkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendLinkListLogic {
	return &FindFriendLinkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取友链列表
func (l *FindFriendLinkListLogic) FindFriendLinkList(in *blog.PageQuery) (*blog.FriendLinkPageResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.FriendLinkModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.FriendLink
	for _, v := range result {
		list = append(list, convert.ConvertFriendLinkModelToPb(v))
	}

	return &blog.FriendLinkPageResp{
		List: list,
	}, nil
}
