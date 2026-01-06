package types

type CommentReviewReq struct {
	Ids      []int64 `json:"ids,optional"`
	IsReview int64   `json:"is_review,optional"`
}

type QueryCommentReq struct {
	PageQuery
	IsReview int64 `json:"is_review,optional"`
	Type     int64 `json:"type,optional"` // 评论类型 1.文章 2.友链 3.说说
}
