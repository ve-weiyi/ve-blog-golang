package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertMenuPb(in *types.MenuDetails) (out *blog.Menu) {
	jsonconv.ObjectToObject(in, &out)

	out.Extra = jsonconv.ObjectToJson(in.Meta)
	return
}

func ConvertMenuTypes(in *blog.Menu) (out *types.MenuDetails) {
	jsonconv.ObjectToObject(in, &out)

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}

func ConvertMenuDetailsTypes(in *blog.MenuDetails) (out *types.MenuDetails) {
	jsonconv.ObjectToObject(in, &out)
	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}
