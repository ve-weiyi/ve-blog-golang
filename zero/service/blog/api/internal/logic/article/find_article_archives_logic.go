package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleArchivesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文章归档(时间轴)
func NewFindArticleArchivesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleArchivesLogic {
	return &FindArticleArchivesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleArchivesLogic) FindArticleArchives(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
