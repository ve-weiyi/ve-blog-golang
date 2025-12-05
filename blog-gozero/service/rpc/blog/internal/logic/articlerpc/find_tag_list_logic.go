package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagListLogic {
	return &FindTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询标签数量
func (l *FindTagListLogic) FindTagList(in *articlerpc.FindTagListReq) (*articlerpc.FindTagListResp, error) {
	helper := NewArticleHelperLogic(l.ctx, l.svcCtx)

	opts := []query.Option{
		query.WithPage(int(in.Paginate.Page)),
		query.WithSize(int(in.Paginate.PageSize)),
		query.WithSorts(in.Paginate.Sorts...),
	}

	if in.TagName != "" {
		opts = append(opts, query.WithCondition("tag_name like ?", "%"+in.TagName+"%"))
	}

	page, size, sorts, conditions, params := query.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TTagModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertTagDetailsResp(records)
	if err != nil {
		return nil, err
	}

	return &articlerpc.FindTagListResp{
		List: list,
		Pagination: &articlerpc.PageResp{
			Page:     in.Paginate.Page,
			PageSize: in.Paginate.PageSize,
			Total:    total,
		},
	}, nil
}
