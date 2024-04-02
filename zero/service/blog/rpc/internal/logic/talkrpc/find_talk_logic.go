package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkLogic {
	return &FindTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询说说
func (l *FindTalkLogic) FindTalk(in *blog.IdReq) (*blog.Talk, error) {
	result, err := l.svcCtx.TalkModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTalkModelToPb(result), nil
}
