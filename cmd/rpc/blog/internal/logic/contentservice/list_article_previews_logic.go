package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type ListArticlePreviewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListArticlePreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListArticlePreviewsLogic {
	return &ListArticlePreviewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询文章预览列表
func (l *ListArticlePreviewsLogic) ListArticlePreviews(in *contentrpc.ListArticlePreviewsRequest) (*contentrpc.ListArticlePreviewsResponse, error) {
	helper := NewArticleHelper(l.ctx, l.svcCtx)
	page, size, sorts, conditions, params, err := helper.convertArticleQuery(toArticleQuery(in))
	if err != nil {
		return nil, err
	}

	records, total, err := l.svcCtx.TArticleModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertArticlePreviewOuts(records)
	if err != nil {
		return nil, err
	}

	return &contentrpc.ListArticlePreviewsResponse{
		ListResult: &contentrpc.ListResult{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
		List: list,
	}, nil
}
