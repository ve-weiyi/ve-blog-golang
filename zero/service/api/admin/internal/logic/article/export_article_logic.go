package article

import (
	"context"
	"path"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导出文章
func NewExportArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportArticleLogic {
	return &ExportArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportArticleLogic) ExportArticle(req *types.IdsReq) (resp *types.EmptyResp, err error) {
	in := &blog.PageQuery{}

	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var aids []int64
	var cids []int64
	for _, v := range out.List {
		aids = append(aids, v.Id)
		cids = append(cids, v.CategoryId)
	}

	// 查询分类
	categories, err := l.svcCtx.CategoryRpc.FindCategoryListByIds(l.ctx, &blog.IdsReq{Ids: cids})
	if err != nil {
		return nil, err
	}
	// 查询标签
	tms, err := l.svcCtx.TagRpc.FindTagMapByArticleIds(l.ctx, &blog.IdsReq{Ids: aids})
	if err != nil {
		return nil, err
	}
	// 转换数据
	for _, v := range out.List {
		var category string
		for _, c := range categories.List {
			if v.CategoryId == c.Id {
				category = c.CategoryName
			}
		}

		var tags []string
		ts := tms.TagMapList[v.Id].List
		if ts != nil {
			for _, t := range ts {
				tags = append(tags, t.TagName)
			}
		}

		m := convert.ConvertArticleHomeTypes(v)
		m.CategoryName = category
		m.TagNameList = tags
		err = l.exportArticle(m)
		if err != nil {
			return nil, err
		}
	}

	return &types.EmptyResp{}, nil
}

func (l *ExportArticleLogic) exportArticle(a *types.ArticleHomeDTO) (err error) {
	fn := path.Join("./runtime/article", a.ArticleTitle+".md")

	ac := invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    fn,
		TemplateString: articleTemplate,
		FunMap:         nil,
		Data: map[string]any{
			"ArticleTitle":    a.ArticleTitle,
			"ArticleCover":    a.ArticleCover,
			"ArticleType":     a.Type,
			"ArticleCategory": a.CategoryName,
			"ArticleTags":     a.TagNameList,
			"ArticleContent":  a.ArticleContent,
		},
	}

	return ac.Execute()
}

const articleTemplate = `
# {{.ArticleTitle}}
文章封面: {{.ArticleCover}}
文章类型: {{.ArticleType}}
文章分类: {{.ArticleCategory}}
文章标签: {{.ArticleTags}}
文章内容:
{{.ArticleContent}}

`
