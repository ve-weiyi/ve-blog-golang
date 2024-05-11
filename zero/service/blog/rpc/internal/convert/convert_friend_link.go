package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertFriendLinkPbToModel(in *blog.FriendLink) (out *model.FriendLink) {
	out = &model.FriendLink{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		// CreatedAt: time.Unix(in.CreatedAt, 0),
		// UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertFriendLinkModelToPb(in *model.FriendLink) (out *blog.FriendLink) {
	out = &blog.FriendLink{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
		CreatedAt:   in.CreatedAt.Unix(),
		UpdatedAt:   in.UpdatedAt.Unix(),
	}

	return out
}
