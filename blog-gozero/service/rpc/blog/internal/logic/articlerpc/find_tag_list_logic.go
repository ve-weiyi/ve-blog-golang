package articlerpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/queryx"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
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

	var opts []queryx.Option
	if in.Paginate != nil {
		opts = append(opts, queryx.WithPage(int(in.Paginate.Page)))
		opts = append(opts, queryx.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, queryx.WithSorts(in.Paginate.Sorts...))
	}

	if in.TagName != "" {
		opts = append(opts, queryx.WithCondition("tag_name like ?", "%"+in.TagName+"%"))
	}

	page, size, sorts, conditions, params := queryx.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TTagModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertTag(records)
	if err != nil {
		return nil, err
	}

	return &articlerpc.FindTagListResp{
		List: list,
		Pagination: &articlerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}
