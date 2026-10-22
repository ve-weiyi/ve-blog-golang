package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/enrich"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/discussionservice"
)

type QueryCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论列表(后台)
func NewQueryCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentListLogic {
	return &QueryCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCommentListLogic) QueryCommentList(req *types.QueryCommentListReq) (resp *types.ListResult, err error) {
	out, err := l.svcCtx.DiscussionService.ListComments(l.ctx, &discussionservice.ListCommentsRequest{
		ListQuery: &discussionservice.ListQuery{Page: req.Page, PageSize: req.PageSize, Sorts: req.Sorts},
		UserId:    req.UserId,
		Status:    req.Status,
		Type:      req.Type,
		TopicId:   req.TopicId,
		ParentId:  req.ParentId,
		IsReview:  req.IsReview,
	})
	if err != nil {
		return nil, err
	}

	var uids, tids []string
	for _, v := range out.List {
		uids = append(uids, v.UserId)
		if v.ReplyUserId != "" {
			uids = append(uids, v.ReplyUserId)
		}
		tids = append(tids, v.DeviceId)
	}

	usm, err := enrich.GetUserInfos(l.ctx, l.svcCtx, uids)
	if err != nil {
		return nil, err
	}

	vsm, err := enrich.GetGuests(l.ctx, l.svcCtx, tids)
	if err != nil {
		return nil, err
	}

	var list []*types.CommentVO
	for _, v := range out.List {
		list = append(list, &types.CommentVO{
			Id:             v.Id,
			UserId:         v.UserId,
			DeviceId:       v.DeviceId,
			Type:           v.Type,
			ReplyUserId:    v.ReplyUserId,
			CommentContent: v.CommentContent,
			Status:         v.Status,
			IsReview:       v.IsReview,
			CreatedAt:      v.CreatedAt,
			UserInfo:       usm[v.UserId],
			GuestInfo:      vsm[v.DeviceId],
			ReplyUserInfo:  usm[v.ReplyUserId],
		})
	}

	return &types.ListResult{
		Page:     out.ListResult.Page,
		PageSize: out.ListResult.PageSize,
		Total:    out.ListResult.Total,
		List:     list,
	}, nil
}
