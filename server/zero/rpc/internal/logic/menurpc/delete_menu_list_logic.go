package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuListLogic {
	return &DeleteMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuListLogic) DeleteMenuList(in *blog.IdsReq) (*blog.BatchResult, error) {
	// todo: add your logic here and delete this line

	return &blog.BatchResult{}, nil
}
