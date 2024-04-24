package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertMenuTypes(in *blog.Menu) (out *types.MenuDetails) {
	jsonconv.ObjectMarshal(in, &out)

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}

func ConvertMenuPb(in *types.MenuDetails) (out *blog.Menu) {
	jsonconv.ObjectMarshal(in, &out)

	out.Extra = jsonconv.ObjectToJson(in.Meta)
	return
}

func ConvertMenuDetailsTypes(in *blog.MenuDetails) (out *types.MenuDetails) {
	jsonconv.ObjectMarshal(in, &out)
	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}
