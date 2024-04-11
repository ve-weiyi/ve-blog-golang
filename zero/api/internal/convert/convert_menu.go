package convert

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/rolerpc"
)

func ConvertMenuTypes(in *rolerpc.Menu) (out *types.Menu) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertMenuPb(in *types.Menu) (out *rolerpc.Menu) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertMenuDetailsTypes(in *rolerpc.MenuDetailsDTO) (out *types.MenuDetailsDTO) {
	err := jsonconv.ObjectMarshal(in, &out)
	if err != nil {
		fmt.Println("err 1--->", err)
	}
	err = jsonconv.JsonToObject(in.Extra, &out.Meta)
	if err != nil {
		fmt.Println("err 2--->", err)
	}
	return
}
