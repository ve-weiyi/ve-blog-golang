package photo_album

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取相册列表
func NewFindPhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumListLogic {
	return &FindPhotoAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumListLogic) FindPhotoAlbumList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.PhotoRpc.FindPhotoAlbumList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.PhotoRpc.FindPhotoAlbumCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.PhotoAlbum
	for _, v := range out.List {
		m := convert.ConvertPhotoAlbumTypes(v)
		count, _ := l.svcCtx.PhotoRpc.FindPhotoCount(l.ctx, &blogrpc.PageQuery{
			Conditions: "album_id = ?",
			Args:       []string{cast.ToString(v.Id)},
		})
		m.PhotoCount = count.Count
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
