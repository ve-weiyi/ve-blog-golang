package article

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/contentservice"
)

type ExportArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导出文章列表
func NewExportArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportArticleLogic {
	return &ExportArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ExportArticle 把选中文章导出为 Markdown；未选文章时导出最近 1000 篇。
func (l *ExportArticleLogic) ExportArticle(req *types.ExportArticleReq) (resp *types.ExportArticleResp, err error) {
	pageSize := len(req.Ids)
	if pageSize == 0 {
		pageSize = 1000
	}

	out, err := l.svcCtx.ContentService.ListArticles(l.ctx, &contentservice.ListArticlesRequest{
		ListQuery: &contentservice.ListQuery{Page: 1, PageSize: int64(pageSize)},
		Ids:       req.Ids,
	})
	if err != nil {
		return nil, err
	}

	var b strings.Builder
	b.WriteString("# 文章导出\n\n")
	b.WriteString(fmt.Sprintf("导出时间：%s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	for _, v := range out.List {
		category := ""
		if v.Category != nil {
			category = v.Category.CategoryName
		}

		b.WriteString("---\n\n")
		b.WriteString("## " + v.ArticleTitle + "\n\n")
		b.WriteString(fmt.Sprintf("- 分类：%s\n", category))
		b.WriteString(fmt.Sprintf("- 状态：%d\n", v.Status))
		b.WriteString(fmt.Sprintf("- 发布时间：%s\n\n", time.UnixMilli(v.CreatedAt).Format("2006-01-02 15:04:05")))
		b.WriteString(v.ArticleContent)
		b.WriteString("\n\n")
	}

	return &types.ExportArticleResp{
		FileName: fmt.Sprintf("articles-%s.md", time.Now().Format("20060102-150405")),
		Content:  b.String(),
	}, nil
}
