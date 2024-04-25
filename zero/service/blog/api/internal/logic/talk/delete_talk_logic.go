package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除说说
func NewDeleteTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkLogic {
	return &DeleteTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTalkLogic) DeleteTalk(req *types.IdReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
