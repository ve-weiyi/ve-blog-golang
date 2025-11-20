/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MigrateConfig struct {
	initFile string
	dataFile string

	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Config   string
}

// migrateCmd represents the migrate command
type MigrateCmd struct {
	cmd *cobra.Command
	cfg *MigrateConfig
}

func NewMigrateCmd() *MigrateCmd {
	migrateCmd := &MigrateCmd{
		cfg: &MigrateConfig{
			initFile: "blog-veweiyi-init.sql",
			dataFile: "blog-veweiyi-data.sql",

			Host:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			Dbname:   "blog-veweiyi",
			Config:   "charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
	migrateCmd.cmd = &cobra.Command{
		Use:   "migrate",
		Short: "初始化数据库表",
		Long:  `初始化数据库表，支持自定义数据库配置和sql文件`,
		Run: func(cmd *cobra.Command, args []string) {
			migrateCmd.RunMigrate(cmd, args)
		},
	}

	migrateCmd.init()
	return migrateCmd
}

func (s *MigrateCmd) init() {
	s.cmd.Flags().StringVarP(&s.cfg.initFile, "file", "i", s.cfg.initFile, "数据库结构sql文件")
	s.cmd.Flags().StringVarP(&s.cfg.dataFile, "data", "d", s.cfg.dataFile, "数据库初始数据sql文件")

	s.cmd.Flags().StringVarP(&s.cfg.Host, "host", "", s.cfg.Host, "数据库ip")
	s.cmd.Flags().StringVarP(&s.cfg.Port, "port", "", s.cfg.Port, "数据库端口")
	s.cmd.Flags().StringVarP(&s.cfg.Username, "username", "", s.cfg.Username, "账号")
	s.cmd.Flags().StringVarP(&s.cfg.Password, "password", "", s.cfg.Password, "密码")
	s.cmd.Flags().StringVarP(&s.cfg.Dbname, "name", "", s.cfg.Dbname, "数据库名称")
	s.cmd.Flags().StringVarP(&s.cfg.Config, "config", "", s.cfg.Config, "数据库配置")
}

func (s *MigrateCmd) RunMigrate(cmd *cobra.Command, args []string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", s.cfg.Username, s.cfg.Password, s.cfg.Host, s.cfg.Port, "", s.cfg.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: "",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
		// gorm日志模式
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connect to ", dsn)

	// 清空数据库
	err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", s.cfg.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	// 创建数据库（如果不存在）
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", s.cfg.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(fmt.Sprintf("USE `%s`", s.cfg.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	// 初始化表
	err = s.ExecSqlFile(db, s.cfg.initFile)
	if err != nil {
		log.Fatal(err)
	}

	// 导入数据
	err = s.ExecSqlFile(db, s.cfg.initFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("init database success")
}

func (s *MigrateCmd) ExecSqlFile(db *gorm.DB, sqlFile string) (err error) {
	if sqlFile == "" {
		return nil
	}

	// 读取 SQL 文件内容
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	// 使用strings.NewReader创建一个读取器，并传递给bufio.NewScanner
	scanner := bufio.NewScanner(strings.NewReader(string(content)))

	// 运行语句
	var query string
	// 按行读取内容
	for scanner.Scan() {
		line := scanner.Text()

		// 忽略注释
		if strings.HasPrefix(line, "--") {
			continue
		}

		// 拼接 SQL 语句
		query += "\n" + line
		if strings.HasSuffix(query, ";") {
			query = strings.TrimSpace(query)
			if err = db.Exec(query).Error; err != nil {
				log.Fatal(err)
			}
			query = ""
		} else {
			continue
		}
	}

	// 检查是否有错误发生（例如：文件是否读取完整）
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
