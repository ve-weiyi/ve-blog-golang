package categoryrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCategoryListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCategoryListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCategoryListByIdsLogic {
	return &FindCategoryListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章分类列表(通过ids)
func (l *FindCategoryListByIdsLogic) FindCategoryListByIds(in *blog.IdsReq) (*blog.CategoryPageResp, error) {
	result, err := l.svcCtx.CategoryModel.FindList(l.ctx, 0, 0, "", "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	var list []*blog.Category
	for _, v := range result {
		list = append(list, convert.ConvertCategoryModelToPb(v))
	}

	return &blog.CategoryPageResp{
		List: list,
	}, nil
}
