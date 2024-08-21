package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

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
func (l *FindTagListLogic) FindTagList(in *blog.FindTagListReq) (*blog.FindTagListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = in.Sorts

	if in.TagName != "" {
		conditions += "tag_name like ?"
		params = append(params, "%"+in.TagName+"%")
	}

	records, err := l.svcCtx.TagModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.TagModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	acm, err := findArticleCountGroupTag(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err
	}

	var list []*blog.TagDetails
	for _, v := range records {
		list = append(list, convertTagOut(v, acm))
	}

	return &blog.FindTagListResp{
		List:  list,
		Total: count,
	}, nil
}
