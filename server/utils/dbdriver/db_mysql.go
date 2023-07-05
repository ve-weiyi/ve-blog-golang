package dbdriver

import (
	"gorm.io/gorm"
)

type MysqlDriver struct {
	*gorm.DB
}

// GetDB 获取数据库的所有数据库名
func (m *MysqlDriver) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME FROM information_schema.schemata;"
	err = m.DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
func (m *MysqlDriver) GetTables(dbName string) (data []Table, err error) {
	var entities []Table
	sql := `select * from information_schema.tables where table_schema = ?`
	err = m.DB.Raw(sql, dbName).Scan(&entities).Error

	return entities, err
}

// GetTableColumns  struct
func (m *MysqlDriver) GetTableColumns(dbName string, tableName string) (data []Column, err error) {
	var entities []Column
	sql := `SELECT * FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_schema = ? AND table_name = ?`
	err = m.DB.Raw(sql, dbName, tableName).Scan(&entities).Error

	return entities, nil
}
