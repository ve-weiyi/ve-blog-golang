package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTalkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkCountLogic {
	return &FindTalkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询说说数量
func (l *FindTalkCountLogic) FindTalkCount(in *blog.PageQuery) (*blog.CountResp, error) {
	_, _, _, conditions, params := convert.ParsePageQuery(in)

	count, err := l.svcCtx.TalkModel.FindCount(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
