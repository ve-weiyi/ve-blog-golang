package response

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Column struct {
	ColumnName    string `json:"columnName" gorm:"column:column_name"`       //列名
	ColumnType    string `json:"columnType" gorm:"column:column_type"`       //字段类型 varchar(11)
	ColumnDefault string `json:"columnDefault" gorm:"column:column_default"` //默认值
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"` //备注
	DataType      string `json:"dataType" gorm:"column:data_type"`           //数据类型 varchar
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`  //数据长度
	IsNullable    string `json:"isNullable" gorm:"column:is_nullable"`       //是否可空
}
