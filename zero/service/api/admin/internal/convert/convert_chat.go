package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
)

func ConvertChatRecordPb(in *types.ChatRecord) (out *blog.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)
	return
}

func ConvertChatRecordTypes(in *blog.ChatRecord) (out *types.ChatRecord) {
	jsonconv.ObjectToObject(in, &out)

	return
}
