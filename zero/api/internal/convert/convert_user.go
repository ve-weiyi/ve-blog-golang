package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/userrpc"
)

func ConvertUserDetailsTypes(in *userrpc.UserDTO) (out *types.UserDTO) {
	jsonconv.ObjectMarshal(in, &out)
	return out
}

func ConvertLoginHistoryTypes(in *userrpc.LoginHistory) (out *types.LoginHistory) {
	jsonconv.ObjectMarshal(in, &out)
	return out
}
