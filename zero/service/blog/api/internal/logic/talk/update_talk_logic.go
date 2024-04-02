package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新说说
func NewUpdateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTalkLogic {
	return &UpdateTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTalkLogic) UpdateTalk(reqCtx *types.RestHeader, req *types.TalkDetails) (resp *types.TalkDetails, err error) {
	in := convert.ConvertTalkPb(req)

	api, err := l.svcCtx.TalkRpc.UpdateTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTalkTypes(api), nil
}
