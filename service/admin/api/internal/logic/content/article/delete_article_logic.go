package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/articleservice"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (resp *types.BatchResp, err error) {
	isDelete := int64(1)
	_, err = l.svcCtx.ArticleService.PatchArticle(l.ctx, &articleservice.PatchArticleRequest{
		Id:       req.Id,
		IsDelete: &isDelete,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: 1,
	}, nil
}
