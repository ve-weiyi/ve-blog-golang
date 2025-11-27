package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkListLogic {
	return &FindTalkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取说说列表
func (l *FindTalkListLogic) FindTalkList(in *talkrpc.FindTalkListReq) (*talkrpc.FindTalkListResp, error) {
	page, size, sorts, conditions, params := convertTalkQuery(in)

	records, total, err := l.svcCtx.TTalkModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*talkrpc.TalkDetails
	for _, v := range records {
		list = append(list, convertTalkOut(v))
	}

	return &talkrpc.FindTalkListResp{
		List: list,
		Pagination: &talkrpc.PageResp{
			Page:     in.Paginate.Page,
			PageSize: in.Paginate.PageSize,
			Total:    total,
		},
	}, nil
}

func convertTalkQuery(in *talkrpc.FindTalkListReq) (page int, size int, sorts string, conditions string, params []any) {
	opts := []query.Option{
		query.WithPage(int(in.Paginate.Page)),
		query.WithSize(int(in.Paginate.PageSize)),
		query.WithSorts(in.Paginate.Sorts...),
	}

	if in.Status != 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}

	return query.NewQueryBuilder(opts...).Build()
}
