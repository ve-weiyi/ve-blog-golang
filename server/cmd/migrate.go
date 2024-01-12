/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
)

// migrateCmd represents the migrate command
type MigrateCmd struct {
	cmd      *cobra.Command
	Mysql    properties.Mysql
	createDB bool
	sqlFile  string
}

func NewMigrateCmd() *MigrateCmd {
	migrateCmd := &MigrateCmd{}
	migrateCmd.cmd = &cobra.Command{
		Use:   "migrate",
		Short: "合并数据库表",
		Long:  `合并数据库表，更新到最新版本`,
		Run: func(cmd *cobra.Command, args []string) {
			migrateCmd.MigrateDB()
		},
	}
	migrateCmd.init()
	return migrateCmd
}

func (s *MigrateCmd) init() {
	s.cmd.PersistentFlags().BoolVarP(&s.createDB, "create", "", false, "是否创建数据库")
	s.cmd.PersistentFlags().StringVarP(&s.sqlFile, "file", "f", "blog-veweiyi.sql", "数据库sql文件")

	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Username, "username", "u", "root", "账号")
	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Password, "password", "p", "123456", "密码")
	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Host, "host", "", "localhost", "数据库ip")
	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Port, "port", "", "3306", "数据库端口")
	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Dbname, "name", "n", "blog", "数据库名称")
	s.cmd.PersistentFlags().StringVarP(&s.Mysql.Config, "config", "c", "charset=utf8mb4&parseTime=True&loc=Local", "数据库配置")
}

func (s *MigrateCmd) MigrateDB() {
	if s.createDB {
		s.InitDatabase()
	}

	dsn := s.Mysql.Dsn()
	fmt.Println("connect to ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 读取 SQL 文件内容
	content, err := os.ReadFile(s.sqlFile)
	if err != nil {
		log.Fatal(err)
	}

	// 将文件内容转为字符串
	queries := string(content)

	// 分割 SQL 命令
	queryList := strings.Split(queries, ";")

	// 执行每个 SQL 命令
	for _, query := range queryList {
		query = strings.TrimSpace(query)
		if query != "" {
			db.Exec(query)
		}
	}

	fmt.Println("Database initialized successfully.")
}

// 创建数据库
func (s *MigrateCmd) InitDatabase() {
	m := s.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, "", m.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建数据库（如果不存在）
	result := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", m.Dbname))
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 关闭连接
	sqlDB, err := db.DB()
	sqlDB.Close()
}
