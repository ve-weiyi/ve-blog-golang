package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindChatRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatRecordsLogic {
	return &FindChatRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindChatRecordsLogic) FindChatRecords(req *types.PageQuery) (resp []types.ChatRecord, err error) {
	// todo: add your logic here and delete this line

	return
}
