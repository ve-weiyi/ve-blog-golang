package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagMapByArticleIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagMapByArticleIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagMapByArticleIdsLogic {
	return &FindTagMapByArticleIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章标签列表
func (l *FindTagMapByArticleIdsLogic) FindTagMapByArticleIds(in *blog.IdsReq) (*blog.TagMapResp, error) {
	ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "article_id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	var tidm = make(map[int64][]int64)
	var tids []int64
	for _, v := range ats {
		tidm[v.ArticleId] = append(tidm[v.ArticleId], v.TagId)
		tids = append(tids, v.TagId)
	}

	result, err := l.svcCtx.TagModel.FindALL(l.ctx, "id in ?", tids)
	if err != nil {
		return nil, err
	}

	var tm = make(map[int64]*blog.TagPageResp)
	for k, v := range tidm {
		var list []*blog.Tag
		for _, vv := range v {
			for _, tag := range result {
				if tag.Id == vv {
					list = append(list, convert.ConvertTagModelToPb(tag))
				}
			}
		}
		tm[k] = &blog.TagPageResp{
			Total: int64(len(list)),
			List:  list,
		}
	}

	return &blog.TagMapResp{
		TagMapList: tm,
	}, nil
}
