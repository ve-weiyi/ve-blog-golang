package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/model"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章详情
func (l *GetArticleLogic) GetArticle(in *contentrpc.GetArticleRequest) (*contentrpc.GetArticleResponse, error) {
	helper := NewArticleHelper(l.ctx, l.svcCtx)

	conditions := "`id` = ?"
	args := []any{in.Id}
	if in.PublicOnly {
		conditions += " AND `status` = ? AND `is_delete` = ?"
		args = append(args, enums.ArticleStatusPublic, 0)
	}
	record, err := l.svcCtx.TArticleModel.FindOne(l.ctx, "", conditions, args...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertArticleOut([]*model.TArticle{record})
	if err != nil {
		return nil, err
	}

	return &contentrpc.GetArticleResponse{
		Article: list[0],
	}, nil
}
