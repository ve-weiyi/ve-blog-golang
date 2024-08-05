package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保存文章
func NewSaveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveArticleLogic {
	return &SaveArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveArticleLogic) SaveArticle(req *types.ArticleNewReq) (resp *types.EmptyResp, err error) {
	in := &blogrpc.Article{}
	jsonconv.ObjectToObject(req, &in)

	c, err := l.svcCtx.CategoryRpc.FindCategoryList(l.ctx, &blogrpc.PageQuery{
		Conditions: "category_name = ?",
		Args:       []string{req.CategoryName},
	})

	// t, err := l.svcCtx.TagRpc.FindTagList(l.ctx, &blogrpc.PageQuery{
	//	Conditions: "tag_name = ?",
	//	Args:       req.TagNameList,
	// })

	if len(c.List) > 0 {
		in.CategoryId = c.List[0].Id
	}

	// if len(t.List)>0 {
	//	for _, tag := range t.List {
	//		in.TagIds = append(in.TagIds, tag.Id)
	//	}
	// }

	_, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
