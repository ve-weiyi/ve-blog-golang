package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
)

func ConvertMenuTypes(in *rolerpc.Menu) (out *types.MenuDetails) {
	jsonconv.ObjectMarshal(in, &out)

	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}

func ConvertMenuPb(in *types.MenuDetails) (out *rolerpc.Menu) {
	jsonconv.ObjectMarshal(in, &out)

	out.Extra = jsonconv.ObjectToJson(in.Meta)
	return
}

func ConvertMenuDetailsTypes(in *rolerpc.MenuDetails) (out *types.MenuDetails) {
	jsonconv.ObjectMarshal(in, &out)
	jsonconv.JsonToObject(in.Extra, &out.Meta)
	return
}
