package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumListLogic {
	return &FindPhotoAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取相册列表
func (l *FindPhotoAlbumListLogic) FindPhotoAlbumList(in *blog.PageQuery) (*blog.FindPhotoAlbumListResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.PhotoAlbumModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.PhotoAlbum
	for _, v := range result {
		list = append(list, convert.ConvertPhotoAlbumModelToPb(v))
	}

	return &blog.FindPhotoAlbumListResp{
		List: list,
	}, nil
}
