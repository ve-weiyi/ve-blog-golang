package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagArticleCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagArticleCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagArticleCountLogic {
	return &FindTagArticleCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询标签关联文章数量
func (l *FindTagArticleCountLogic) FindTagArticleCount(in *blog.FindTagArticleCountReq) (*blog.CountResp, error) {
	count, err := l.svcCtx.ArticleTagModel.FindCount(l.ctx, "tag_id = ?", in.TagId)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
