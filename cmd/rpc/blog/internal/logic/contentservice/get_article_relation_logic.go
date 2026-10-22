package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

type GetArticleRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleRelationLogic {
	return &GetArticleRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询关联文章
func (l *GetArticleRelationLogic) GetArticleRelation(in *contentrpc.GetArticleRelationRequest) (*contentrpc.GetArticleRelationResponse, error) {
	helper := NewArticleHelper(l.ctx, l.svcCtx)

	record, err := l.svcCtx.TArticleModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	last, _, err := l.svcCtx.TArticleModel.FindListAndTotal(l.ctx, 1, 1, "id desc", "id < ? and status = ? and is_delete = 0", record.Id, enums.ArticleStatusPublic)
	if err != nil {
		return nil, err
	}
	next, _, err := l.svcCtx.TArticleModel.FindListAndTotal(l.ctx, 1, 1, "id asc", "id > ? and status = ? and is_delete = 0", record.Id, enums.ArticleStatusPublic)
	if err != nil {
		return nil, err
	}
	recommend, _, err := l.svcCtx.TArticleModel.FindListAndTotal(l.ctx, 1, 5, "id asc", "id != ? and category_id = ? and status = ? and is_delete = 0", record.Id, record.CategoryId, enums.ArticleStatusPublic)
	if err != nil {
		return nil, err
	}
	newest, _, err := l.svcCtx.TArticleModel.FindListAndTotal(l.ctx, 1, 5, "id desc", "status = ? and is_delete = 0", enums.ArticleStatusPublic)
	if err != nil {
		return nil, err
	}

	resp := &contentrpc.GetArticleRelationResponse{}
	if len(last) > 0 {
		resp.Last = helper.convertArticlePreviewOut(last[0])
	}
	if len(next) > 0 {
		resp.Next = helper.convertArticlePreviewOut(next[0])
	}
	for _, v := range recommend {
		resp.Recommends = append(resp.Recommends, helper.convertArticlePreviewOut(v))
	}
	for _, v := range newest {
		resp.Latest = append(resp.Latest, helper.convertArticlePreviewOut(v))
	}

	return resp, nil
}
