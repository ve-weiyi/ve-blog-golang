package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertPhotoTypes(in *blog.Photo) (out *types.Photo) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertPhotoPb(in *types.Photo) (out *blog.Photo) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
