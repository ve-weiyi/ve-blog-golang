package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertPhotoTypes(in *blog.Photo) (out *types.Photo) {
	jsonconv.ObjectToObject(in, &out)

	return
}

func ConvertPhotoAlbumTypes(in *blog.PhotoAlbum) (out *types.Album) {
	jsonconv.ObjectToObject(in, &out)

	return
}
