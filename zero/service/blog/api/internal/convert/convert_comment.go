package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertCommentTypes(in *blog.Comment) (out *types.Comment) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertCommentPb(in *types.Comment) (out *blog.Comment) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertCommentBackTypes(in *blog.Comment) (out *types.CommentBackDTO) {
	out = &types.CommentBackDTO{
		Id:             in.Id,
		Type:           in.Type,
		TopicTitle:     "",
		UserAvatar:     "",
		UserNickname:   "",
		CommentContent: in.CommentContent,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt,
	}

	return
}
