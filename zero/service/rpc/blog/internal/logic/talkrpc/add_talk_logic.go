package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTalkLogic {
	return &AddTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建说说
func (l *AddTalkLogic) AddTalk(in *blog.Talk) (*blog.Talk, error) {
	entity := convert.ConvertTalkPbToModel(in)

	_, err := l.svcCtx.TalkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTalkModelToPb(entity), nil
}
