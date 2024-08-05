package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertPhotoPb(in *types.Photo) (out *blogrpc.Photo) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertPhotoTypes(in *blogrpc.Photo) (out *types.Photo) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertPhotoAlbumTypes(in *blogrpc.PhotoAlbum) (out *types.PhotoAlbum) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertPhotoAlbumPb(in *types.PhotoAlbum) (out *blogrpc.PhotoAlbum) {
	jsonconv.ObjectToObject(in, &out)
	return
}
