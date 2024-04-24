package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleListLogic {
	return &DeleteArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleListLogic) DeleteArticleList(in *blog.IdsReq) (*blog.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &blog.BatchResp{}, nil
}
