package enums

// 通知消息状态（与 DB 存储值一致）
const (
	NoticeStatusDraft     = "draft"
	NoticeStatusPublished = "published"
	NoticeStatusRevoked   = "revoked"
)

// 通知渠道
const (
	NotifyChannelInbox = "inbox"
	NotifyChannelEmail = "email"
	NotifyChannelSms   = "sms"
)

// 投递记录状态
const (
	NotifyRecordStatusUnread  = "unread"
	NotifyRecordStatusRead    = "read"
	NotifyRecordStatusRevoked = "revoked"
	NotifyRecordStatusPending = "pending"
	NotifyRecordStatusSent    = "sent"
	NotifyRecordStatusFailed  = "failed"
)

// 通知目标类型
const (
	NotifyTargetTypeAll     = "all"
	NotifyTargetTypeUserIDs = "user_ids"
)
