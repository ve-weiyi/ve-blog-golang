package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertChatRecordPb(in *types.ChatRecord) (out *blog.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertChatRecordTypes(in *blog.ChatRecord) (out *types.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)

	return
}
