package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/photorpc"
)

func ConvertPhotoTypes(in *photorpc.PhotoNew) (out *types.Photo) {

	return &types.Photo{
		Id:       in.Id,
		PhotoUrl: in.PhotoSrc,
	}
}

func ConvertPhotoAlbumTypes(in *photorpc.AlbumDetails) (out *types.Album) {

	return &types.Album{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
	}
}
