package entity

import "time"

// TableNamePhoto return the table name of <photo>
const TableNamePhoto = "photo"

// Photo mapped from table <photo>
type Photo struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 主键
	AlbumId   int64     `gorm:"column:album_id" json:"album_id" `     // 相册id
	PhotoName string    `gorm:"column:photo_name" json:"photo_name" ` // 照片名
	PhotoDesc string    `gorm:"column:photo_desc" json:"photo_desc" ` // 照片描述
	PhotoSrc  string    `gorm:"column:photo_src" json:"photo_src" `   // 照片地址
	IsDelete  int64     `gorm:"column:is_delete" json:"is_delete" `   // 是否删除
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Photo 's table name
func (*Photo) TableName() string {
	return TableNamePhoto
}
