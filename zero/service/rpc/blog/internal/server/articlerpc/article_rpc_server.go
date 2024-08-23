// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type ArticleRpcServer struct {
	svcCtx *svc.ServiceContext
	articlerpc.UnimplementedArticleRpcServer
}

func NewArticleRpcServer(svcCtx *svc.ServiceContext) *ArticleRpcServer {
	return &ArticleRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建文章
func (s *ArticleRpcServer) AddArticle(ctx context.Context, in *articlerpc.ArticleNew) (*articlerpc.ArticleDetails, error) {
	l := articlerpclogic.NewAddArticleLogic(ctx, s.svcCtx)
	return l.AddArticle(in)
}

// 更新文章
func (s *ArticleRpcServer) UpdateArticle(ctx context.Context, in *articlerpc.ArticleNew) (*articlerpc.ArticleDetails, error) {
	l := articlerpclogic.NewUpdateArticleLogic(ctx, s.svcCtx)
	return l.UpdateArticle(in)
}

// 查询文章
func (s *ArticleRpcServer) GetArticle(ctx context.Context, in *articlerpc.IdReq) (*articlerpc.ArticleDetails, error) {
	l := articlerpclogic.NewGetArticleLogic(ctx, s.svcCtx)
	return l.GetArticle(in)
}

// 删除文章
func (s *ArticleRpcServer) DeleteArticle(ctx context.Context, in *articlerpc.IdsReq) (*articlerpc.BatchResp, error) {
	l := articlerpclogic.NewDeleteArticleLogic(ctx, s.svcCtx)
	return l.DeleteArticle(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticleList(ctx context.Context, in *articlerpc.FindArticleListReq) (*articlerpc.FindArticleListResp, error) {
	l := articlerpclogic.NewFindArticleListLogic(ctx, s.svcCtx)
	return l.FindArticleList(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticlePublicList(ctx context.Context, in *articlerpc.FindArticleListReq) (*articlerpc.FindArticleListResp, error) {
	l := articlerpclogic.NewFindArticlePublicListLogic(ctx, s.svcCtx)
	return l.FindArticlePublicList(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticlesByTag(ctx context.Context, in *articlerpc.FindArticlesByTagReq) (*articlerpc.FindArticleListResp, error) {
	l := articlerpclogic.NewFindArticlesByTagLogic(ctx, s.svcCtx)
	return l.FindArticlesByTag(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticlesByCategory(ctx context.Context, in *articlerpc.FindArticlesByCategoryReq) (*articlerpc.FindArticleListResp, error) {
	l := articlerpclogic.NewFindArticlesByCategoryLogic(ctx, s.svcCtx)
	return l.FindArticlesByCategory(in)
}

// 回收文章
func (s *ArticleRpcServer) RecycleArticle(ctx context.Context, in *articlerpc.RecycleArticleReq) (*articlerpc.EmptyResp, error) {
	l := articlerpclogic.NewRecycleArticleLogic(ctx, s.svcCtx)
	return l.RecycleArticle(in)
}

// 置顶文章
func (s *ArticleRpcServer) TopArticle(ctx context.Context, in *articlerpc.TopArticleReq) (*articlerpc.EmptyResp, error) {
	l := articlerpclogic.NewTopArticleLogic(ctx, s.svcCtx)
	return l.TopArticle(in)
}

// 点赞文章
func (s *ArticleRpcServer) LikeArticle(ctx context.Context, in *articlerpc.IdReq) (*articlerpc.EmptyResp, error) {
	l := articlerpclogic.NewLikeArticleLogic(ctx, s.svcCtx)
	return l.LikeArticle(in)
}

// 用户点赞的文章
func (s *ArticleRpcServer) FindUserLikeArticle(ctx context.Context, in *articlerpc.UserIdReq) (*articlerpc.FindLikeArticleResp, error) {
	l := articlerpclogic.NewFindUserLikeArticleLogic(ctx, s.svcCtx)
	return l.FindUserLikeArticle(in)
}

// 查询文章数量
func (s *ArticleRpcServer) AnalysisArticle(ctx context.Context, in *articlerpc.EmptyReq) (*articlerpc.AnalysisArticleResp, error) {
	l := articlerpclogic.NewAnalysisArticleLogic(ctx, s.svcCtx)
	return l.AnalysisArticle(in)
}

// 创建文章分类
func (s *ArticleRpcServer) AddCategory(ctx context.Context, in *articlerpc.CategoryNew) (*articlerpc.CategoryDetails, error) {
	l := articlerpclogic.NewAddCategoryLogic(ctx, s.svcCtx)
	return l.AddCategory(in)
}

// 更新文章分类
func (s *ArticleRpcServer) UpdateCategory(ctx context.Context, in *articlerpc.CategoryNew) (*articlerpc.CategoryDetails, error) {
	l := articlerpclogic.NewUpdateCategoryLogic(ctx, s.svcCtx)
	return l.UpdateCategory(in)
}

// 查询文章分类
func (s *ArticleRpcServer) GetCategory(ctx context.Context, in *articlerpc.IdReq) (*articlerpc.CategoryDetails, error) {
	l := articlerpclogic.NewGetCategoryLogic(ctx, s.svcCtx)
	return l.GetCategory(in)
}

// 删除文章分类
func (s *ArticleRpcServer) DeleteCategory(ctx context.Context, in *articlerpc.IdsReq) (*articlerpc.BatchResp, error) {
	l := articlerpclogic.NewDeleteCategoryLogic(ctx, s.svcCtx)
	return l.DeleteCategory(in)
}

// 查询文章分类列表
func (s *ArticleRpcServer) FindCategoryList(ctx context.Context, in *articlerpc.FindCategoryListReq) (*articlerpc.FindCategoryListResp, error) {
	l := articlerpclogic.NewFindCategoryListLogic(ctx, s.svcCtx)
	return l.FindCategoryList(in)
}

// 创建标签
func (s *ArticleRpcServer) AddTag(ctx context.Context, in *articlerpc.TagNew) (*articlerpc.TagDetails, error) {
	l := articlerpclogic.NewAddTagLogic(ctx, s.svcCtx)
	return l.AddTag(in)
}

// 更新标签
func (s *ArticleRpcServer) UpdateTag(ctx context.Context, in *articlerpc.TagNew) (*articlerpc.TagDetails, error) {
	l := articlerpclogic.NewUpdateTagLogic(ctx, s.svcCtx)
	return l.UpdateTag(in)
}

// 查询标签
func (s *ArticleRpcServer) GetTag(ctx context.Context, in *articlerpc.IdReq) (*articlerpc.TagDetails, error) {
	l := articlerpclogic.NewGetTagLogic(ctx, s.svcCtx)
	return l.GetTag(in)
}

// 删除标签
func (s *ArticleRpcServer) DeleteTag(ctx context.Context, in *articlerpc.IdsReq) (*articlerpc.BatchResp, error) {
	l := articlerpclogic.NewDeleteTagLogic(ctx, s.svcCtx)
	return l.DeleteTag(in)
}

// 查询标签列表
func (s *ArticleRpcServer) FindTagList(ctx context.Context, in *articlerpc.FindTagListReq) (*articlerpc.FindTagListResp, error) {
	l := articlerpclogic.NewFindTagListLogic(ctx, s.svcCtx)
	return l.FindTagList(in)
}
