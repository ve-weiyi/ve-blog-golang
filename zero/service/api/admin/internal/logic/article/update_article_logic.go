package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保存文章
func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.ArticleNewReq) (resp *types.EmptyResp, err error) {
	in := &articlerpc.ArticleNew{
		Id:             req.Id,
		UserId:         l.ctx.Value("userId").(int64),
		ArticleCover:   req.ArticleCover,
		ArticleTitle:   req.ArticleTitle,
		ArticleContent: req.ArticleContent,
		ArticleType:    req.ArticleType,
		OriginalUrl:    req.OriginalUrl,
		Status:         req.Status,
		CategoryName:   req.CategoryName,
		TagNameList:    req.TagNameList,
	}

	_, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}