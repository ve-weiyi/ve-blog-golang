package dbdriver

type Table struct {
	TableName    string `json:"tableName" gorm:"column:TABLE_NAME"`
	TableComment string `json:"tableComment" gorm:"column:TABLE_COMMENT"`
}
