package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/router"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

// register your handlers here
func RegisterHandlers(r *gin.RouterGroup, svCtx *svctx.ServiceContext) {
	router.NewAlbumRouter(svCtx).Register(r)
	router.NewArticleRouter(svCtx).Register(r)
	router.NewAuthRouter(svCtx).Register(r)
	router.NewCategoryRouter(svCtx).Register(r)
	router.NewCommentRouter(svCtx).Register(r)
	router.NewCommonRouter(svCtx).Register(r)
	router.NewFriendRouter(svCtx).Register(r)
	router.NewPageRouter(svCtx).Register(r)
	router.NewRemarkRouter(svCtx).Register(r)
	router.NewTagRouter(svCtx).Register(r)
	router.NewTalkRouter(svCtx).Register(r)
	router.NewUploadRouter(svCtx).Register(r)
	router.NewUserRouter(svCtx).Register(r)
	router.NewWebsiteRouter(svCtx).Register(r)
	router.NewWebsocketRouter(svCtx).Register(r)
}
