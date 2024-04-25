package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertChatRecordTypes(in *blog.ChatRecord) (out *types.ChatRecord) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertChatRecordPb(in *types.ChatRecord) (out *blog.ChatRecord) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
