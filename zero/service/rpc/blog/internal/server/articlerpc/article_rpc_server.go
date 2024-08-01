// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

type ArticleRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedArticleRpcServer
}

func NewArticleRpcServer(svcCtx *svc.ServiceContext) *ArticleRpcServer {
	return &ArticleRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建文章
func (s *ArticleRpcServer) AddArticle(ctx context.Context, in *blog.Article) (*blog.Article, error) {
	l := articlerpclogic.NewAddArticleLogic(ctx, s.svcCtx)
	return l.AddArticle(in)
}

// 更新文章
func (s *ArticleRpcServer) UpdateArticle(ctx context.Context, in *blog.Article) (*blog.Article, error) {
	l := articlerpclogic.NewUpdateArticleLogic(ctx, s.svcCtx)
	return l.UpdateArticle(in)
}

// 删除文章
func (s *ArticleRpcServer) DeleteArticle(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := articlerpclogic.NewDeleteArticleLogic(ctx, s.svcCtx)
	return l.DeleteArticle(in)
}

// 批量删除文章
func (s *ArticleRpcServer) DeleteArticleList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := articlerpclogic.NewDeleteArticleListLogic(ctx, s.svcCtx)
	return l.DeleteArticleList(in)
}

// 查询文章
func (s *ArticleRpcServer) FindArticle(ctx context.Context, in *blog.IdReq) (*blog.Article, error) {
	l := articlerpclogic.NewFindArticleLogic(ctx, s.svcCtx)
	return l.FindArticle(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticleList(ctx context.Context, in *blog.PageQuery) (*blog.ArticlePageResp, error) {
	l := articlerpclogic.NewFindArticleListLogic(ctx, s.svcCtx)
	return l.FindArticleList(in)
}

// 查询文章数量
func (s *ArticleRpcServer) FindArticleCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := articlerpclogic.NewFindArticleCountLogic(ctx, s.svcCtx)
	return l.FindArticleCount(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticleByTag(ctx context.Context, in *blog.FindArticleByTagReq) (*blog.ArticlePageResp, error) {
	l := articlerpclogic.NewFindArticleByTagLogic(ctx, s.svcCtx)
	return l.FindArticleByTag(in)
}

// 查询文章列表
func (s *ArticleRpcServer) FindArticleByCategory(ctx context.Context, in *blog.FindArticleByCategoryReq) (*blog.ArticlePageResp, error) {
	l := articlerpclogic.NewFindArticleByCategoryLogic(ctx, s.svcCtx)
	return l.FindArticleByCategory(in)
}

// 点赞文章
func (s *ArticleRpcServer) LikeArticle(ctx context.Context, in *blog.IdReq) (*blog.EmptyResp, error) {
	l := articlerpclogic.NewLikeArticleLogic(ctx, s.svcCtx)
	return l.LikeArticle(in)
}

// 用户点赞的文章
func (s *ArticleRpcServer) FindUserLikeArticle(ctx context.Context, in *blog.UserIdReq) (*blog.FindLikeArticleResp, error) {
	l := articlerpclogic.NewFindUserLikeArticleLogic(ctx, s.svcCtx)
	return l.FindUserLikeArticle(in)
}
