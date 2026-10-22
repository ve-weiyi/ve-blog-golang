package enums

// 评论相关枚举 (t_comment)
const (
	// 评论类型 (type)
	CommentTypeArticle = 1 // 文章评论
	CommentTypeFriend  = 2 // 友链评论
	CommentTypeTalk    = 3 // 说说评论

	// 评论状态 (status)
	CommentStatusNormal  = 0 // 正常
	CommentStatusEdited  = 1 // 已编辑
	CommentStatusDeleted = 2 // 已删除

	// 评论审核状态 (is_review)
	CommentIsReviewNo  = 0 // 未审核通过
	CommentIsReviewYes = 1 // 已审核通过
)

// 聊天相关枚举 (t_chat)
const (
	// 聊天消息状态 (status)
	ChatStatusNormal  = 0 // 正常
	ChatStatusEdited  = 1 // 编辑
	ChatStatusRecall  = 2 // 撤回
	ChatStatusDeleted = 3 // 删除
)

// 留言相关枚举 (t_message)
const (
	// 留言状态 (status)
	MessageStatusNormal  = 0 // 正常
	MessageStatusEdited  = 1 // 编辑
	MessageStatusRecall  = 2 // 撤回
	MessageStatusDeleted = 3 // 删除

	// 留言审核状态 (is_review)
	MessageIsReviewNo  = 0 // 未审核通过
	MessageIsReviewYes = 1 // 已审核通过
)

// 说说相关枚举 (t_talk)
const (
	// 说说置顶状态 (is_top)
	TalkIsTopNo  = 0 // 未置顶
	TalkIsTopYes = 1 // 已置顶

	// 说说状态 (status)
	TalkStatusPublic  = 1 // 公开
	TalkStatusPrivate = 2 // 私密
)
