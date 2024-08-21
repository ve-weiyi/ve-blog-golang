package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建文章
func (l *AddArticleLogic) AddArticle(in *blog.ArticleNew) (*blog.ArticleDetails, error) {
	entity := convertArticleIn(in)

	// 插入文章分类
	categoryId, err := findOrAddCategory(l.ctx, l.svcCtx, in.CategoryName)
	if err != nil {
		return nil, err
	}

	entity.CategoryId = categoryId
	_, err = l.svcCtx.ArticleModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	// 插入文章标签
	var ats []*model.ArticleTag
	for _, tagName := range in.TagNameList {
		tagId, err := findOrAddTag(l.ctx, l.svcCtx, tagName)
		if err != nil {
			return nil, err
		}

		at := &model.ArticleTag{
			ArticleId: entity.Id,
			TagId:     tagId,
		}
		ats = append(ats, at)
	}
	l.svcCtx.ArticleTagModel.DeleteBatch(l.ctx, "article_id = ?", entity.Id)
	l.svcCtx.ArticleTagModel.InsertBatch(l.ctx, ats...)

	return &blog.ArticleDetails{}, nil
}
