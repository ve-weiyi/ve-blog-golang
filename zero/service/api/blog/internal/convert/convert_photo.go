package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertPhotoTypes(in *blogrpc.Photo) (out *types.Photo) {

	return &types.Photo{
		Id:       in.Id,
		PhotoUrl: in.PhotoSrc,
	}
}

func ConvertPhotoAlbumTypes(in *blogrpc.PhotoAlbum) (out *types.Album) {

	return &types.Album{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
	}
}
