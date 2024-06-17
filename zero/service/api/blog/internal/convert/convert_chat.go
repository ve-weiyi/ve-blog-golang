package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

func ConvertChatRecordPb(in *types.ChatRecord) (out *blogrpc.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertChatRecordTypes(in *blogrpc.ChatRecord) (out *types.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)

	return
}
