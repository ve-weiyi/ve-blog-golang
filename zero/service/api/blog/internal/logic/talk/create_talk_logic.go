package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建说说
func NewCreateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTalkLogic {
	return &CreateTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTalkLogic) CreateTalk(req *types.TalkDetails) (resp *types.TalkDetails, err error) {
	in := convert.ConvertTalkPb(req)
	out, err := l.svcCtx.TalkRpc.CreateTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertTalkTypes(out)
	return resp, nil
}
