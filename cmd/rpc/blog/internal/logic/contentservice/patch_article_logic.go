package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
)

type PatchArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchArticleLogic {
	return &PatchArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部分更新文章
func (l *PatchArticleLogic) PatchArticle(in *contentrpc.PatchArticleRequest) (*contentrpc.PatchArticleResponse, error) {
	if in.IsDelete != nil && !enums.IsValidIsDelete(*in.IsDelete) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "删除标记不合法")
	}
	if in.IsTop != nil && !enums.IsValidIsDelete(*in.IsTop) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "置顶标记不合法")
	}

	record, err := l.svcCtx.TArticleModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if in.IsDelete != nil {
		record.IsDelete = *in.IsDelete
	}
	if in.IsTop != nil {
		record.IsTop = *in.IsTop
	}

	_, err = l.svcCtx.TArticleModel.Update(l.ctx, record)
	if err != nil {
		return nil, err
	}

	return &contentrpc.PatchArticleResponse{Success: true}, nil
}
