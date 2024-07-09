package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertPhotoPb(in *types.Photo) (out *blog.Photo) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertPhotoTypes(in *blog.Photo) (out *types.Photo) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertPhotoAlbumTypes(in *blog.PhotoAlbum) (out *types.PhotoAlbum) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertPhotoAlbumPb(in *types.PhotoAlbum) (out *blog.PhotoAlbum) {
	jsonconv.ObjectToObject(in, &out)
	return
}
