package request

type CommentQueryReq struct {
	Page     int64  `json:"page,optional"`      // 页码
	PageSize int64  `json:"page_size,optional"` // 每页数量
	TopicId  int64  `json:"topic_id,optional"`  // 主题id
	ParentId int64  `json:"parent_id,optional"` // 父评论id
	Type     int64  `json:"type,optional"`      // 评论类型 1.文章 2.友链 3.说说
	OrderBy  string `json:"order_by,optional"`  // 排序字段 create_at|like_count
}
