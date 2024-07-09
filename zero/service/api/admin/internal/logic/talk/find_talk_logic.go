package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
