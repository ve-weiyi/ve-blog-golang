package model

import (
	"gorm.io/gorm"
)

// 获取数据库信息
func GetSchema(t *gorm.DB) (data *Schema, err error) {
	var models []*Table
	tables, err := t.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	for _, table := range tables {
		m, err := GetTable(t, table)
		if err != nil {
			return nil, err
		}
		models = append(models, m)
	}

	data = &Schema{
		SchemaName: t.Migrator().CurrentDatabase(),
		Tables:     models,
	}

	return data, nil
}

// 获取数据库表信息
func GetTable(t *gorm.DB, tableName string) (data *Table, err error) {
	types, err := t.Migrator().TableType(tableName)
	if err != nil {
		return nil, err
	}

	var out Table
	out.SchemaName = t.Migrator().CurrentDatabase()
	out.TableType = types
	out.TableName = types.Name()
	out.TableComment, _ = types.Comment()
	out.Type = types.Type()
	out.Columns, err = GetTableColumns(t, tableName)
	out.Indexes, err = GetTableIndexes(t, tableName)
	return &out, nil
}

// 获取表字段信息
func GetTableColumns(t *gorm.DB, tableName string) (data []*Column, err error) {
	var entities []*Column

	var mapIndex map[string][]*Index

	types, err := t.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}

	indexes, err := t.Migrator().GetIndexes(tableName)
	if err != nil {
		return nil, err
	}

	mapIndex = GroupByColumn(indexes)
	for _, entity := range types {
		col := &Column{}
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

// 获取表索引信息
func GetTableIndexes(t *gorm.DB, tableName string) (data map[string][]*Index, err error) {
	indexes, err := t.Migrator().GetIndexes(tableName)
	if err != nil {
		return nil, err
	}

	return GroupByColumn(indexes), nil
}

// GroupByColumn group columns
func GroupByColumn(indexList []gorm.Index) map[string][]*Index {
	columnIndexMap := make(map[string][]*Index, len(indexList))
	if len(indexList) == 0 {
		return columnIndexMap
	}

	for _, idx := range indexList {
		if idx == nil {
			continue
		}
		for i, col := range idx.Columns() {
			columnIndexMap[col] = append(columnIndexMap[col], &Index{
				Index:    idx,
				Priority: int32(i + 1),
			})
		}
	}
	return columnIndexMap
}
