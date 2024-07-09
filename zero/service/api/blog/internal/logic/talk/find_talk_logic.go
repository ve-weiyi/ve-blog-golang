package talk

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

type FindTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询说说
func NewFindTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkLogic {
	return &FindTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTalkLogic) FindTalk(req *types.IdReq) (resp *types.TalkDetails, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.TalkRpc.FindTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTalkTypes(out), nil
}
