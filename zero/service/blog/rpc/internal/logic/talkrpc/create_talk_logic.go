package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTalkLogic {
	return &CreateTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建说说
func (l *CreateTalkLogic) CreateTalk(in *blog.Talk) (*blog.Talk, error) {
	entity := convert.ConvertTalkPbToModel(in)

	result, err := l.svcCtx.TalkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTalkModelToPb(result), nil
}
