package entity

import "time"

// TableNameUploadRecord return the table name of <upload_record>
const TableNameUploadRecord = "upload_record"

// UploadRecord mapped from table <upload_record>
type UploadRecord struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // id
	UserId    int64     `gorm:"column:user_id" json:"user_id" `       // 用户id
	Label     string    `gorm:"column:label" json:"label" `           // 标签
	FileName  string    `gorm:"column:file_name" json:"file_name" `   // 文件名称
	FileSize  int64     `gorm:"column:file_size" json:"file_size" `   // 文件大小
	FileMd5   string    `gorm:"column:file_md5" json:"file_md5" `     // 文件md5值
	FileUrl   string    `gorm:"column:file_url" json:"file_url" `     // 上传路径
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName UploadRecord 's table name
func (*UploadRecord) TableName() string {
	return TableNameUploadRecord
}
