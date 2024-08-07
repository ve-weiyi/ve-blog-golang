package entity

import "time"

// TableNamePhotoAlbum return the table name of <photo_album>
const TableNamePhotoAlbum = "photo_album"

// PhotoAlbum mapped from table <photo_album>
type PhotoAlbum struct {
	Id         int64     `gorm:"column:id" json:"id" `                   // 主键
	AlbumName  string    `gorm:"column:album_name" json:"album_name" `   // 相册名
	AlbumDesc  string    `gorm:"column:album_desc" json:"album_desc" `   // 相册描述
	AlbumCover string    `gorm:"column:album_cover" json:"album_cover" ` // 相册封面
	IsDelete   int64     `gorm:"column:is_delete" json:"is_delete" `     // 是否删除
	Status     int64     `gorm:"column:status" json:"status" `           // 状态值 1公开 2私密
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at" `   // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at" `   // 更新时间
}

// TableName PhotoAlbum 's table name
func (*PhotoAlbum) TableName() string {
	return TableNamePhotoAlbum
}
