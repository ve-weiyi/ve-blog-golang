package provider

import (
	"gorm.io/gorm"
)

// GetTableColumns  struct
func GetTableColumnTypes(t *gorm.DB, tableName string) (result []gorm.ColumnType, err error) {
	types, err := t.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// GetTableIndex  index
func GetTableIndex(t *gorm.DB, tableName string) (indexes []gorm.Index, err error) {
	return t.Migrator().GetIndexes(tableName)
}

// GetTableColumns  struct
func GetTableColumns(t *gorm.DB, tableName string) (data []Column, err error) {
	var entities []Column
	//var metas []ColumnMetadata
	//sql := `SELECT * FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_schema = ? AND table_name = ?`
	//err = t.DB.Raw(sql,dbName, tableName).Scan(&metas).Error
	var mapType map[string]gorm.ColumnType
	var mapIndex map[string][]*Index

	types, err := t.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}
	mapType = make(map[string]gorm.ColumnType, 0)
	for _, item := range types {
		mapType[item.Name()] = item
	}

	indexes, err := t.Migrator().GetIndexes(tableName)

	if err != nil {
		return nil, err
	}
	mapIndex = GroupByColumn(indexes)
	for _, entity := range types {
		col := Column{}
		col.ColumnType = entity

		col.ColumnName = entity.Name()
		col.Indexes = mapIndex[entity.Name()]
		col.ColumnFiledType, _ = entity.ColumnType()
		col.ColumnDefault, col.HasDefault = entity.DefaultValue()
		col.ColumnComment, _ = entity.Comment()
		col.DataType = entity.DatabaseTypeName()
		col.DataTypeLong, _ = entity.Length()
		col.IsNullable, _ = entity.Nullable()
		col.IsPrimaryKey, _ = entity.PrimaryKey()
		col.IsUnique, _ = entity.Unique()
		col.IsAutoIncrement, _ = entity.AutoIncrement()
		entities = append(entities, col)
	}

	return entities, nil
}
