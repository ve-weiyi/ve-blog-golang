package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagListByArticleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagListByArticleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagListByArticleIdLogic {
	return &FindTagListByArticleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章标签列表
func (l *FindTagListByArticleIdLogic) FindTagListByArticleId(in *blog.IdReq) (*blog.TagPageResp, error) {
	ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "article_id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	var tids []int64
	for _, v := range ats {
		tids = append(tids, v.TagId)
	}

	result, err := l.svcCtx.TagModel.FindALL(l.ctx, "id in ?", tids)
	if err != nil {
		return nil, err
	}

	var list []*blog.Tag
	for _, v := range result {
		list = append(list, convert.ConvertTagModelToPb(v))
	}

	return &blog.TagPageResp{
		List: list,
	}, nil
}
