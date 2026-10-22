package discussionservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/discussionrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

type GetCommentReplyCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyCountLogic {
	return &GetCommentReplyCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询评论回复数量
func (l *GetCommentReplyCountLogic) GetCommentReplyCount(in *discussionrpc.GetCommentReplyCountRequest) (*discussionrpc.GetCommentReplyCountResponse, error) {
	result := make(map[int64]int64, len(in.TopicIds))
	if len(in.TopicIds) == 0 {
		return &discussionrpc.GetCommentReplyCountResponse{TopicCommentCounts: result}, nil
	}

	var rows []struct {
		TopicId int64 `gorm:"column:topic_id"`
		Count   int64 `gorm:"column:comment_count"`
	}
	err := l.svcCtx.GormDB.Table("`t_comment`").
		Select("topic_id, COUNT(*) AS comment_count").
		Where("topic_id IN ? AND is_review = ? AND status != ?",
			in.TopicIds, enums.CommentIsReviewYes, enums.CommentStatusDeleted).
		Group("topic_id").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		result[row.TopicId] = row.Count
	}
	for _, topicId := range in.TopicIds {
		if _, ok := result[topicId]; !ok {
			result[topicId] = 0
		}
	}

	return &discussionrpc.GetCommentReplyCountResponse{TopicCommentCounts: result}, nil
}
