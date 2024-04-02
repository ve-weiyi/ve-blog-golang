package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTalkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkListLogic {
	return &DeleteTalkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTalkListLogic) DeleteTalkList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
