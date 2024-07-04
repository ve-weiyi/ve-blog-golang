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

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

// migrateCmd represents the migrate command
type MigrateCmd struct {
	cmd     *cobra.Command
	action  string
	sqlFile string

	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Config   string
}

func NewMigrateCmd() *MigrateCmd {
	migrateCmd := &MigrateCmd{}
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
	s.cmd.Flags().StringVarP(&s.action, "action", "a", "migrate", "migrate|reset")
	s.cmd.Flags().StringVarP(&s.sqlFile, "file", "", "blog-veweiyi.sql", "数据库sql文件")

	s.cmd.Flags().StringVarP(&s.Host, "host", "", "localhost", "数据库ip")
	s.cmd.Flags().StringVarP(&s.Port, "port", "", "3306", "数据库端口")
	s.cmd.Flags().StringVarP(&s.Username, "username", "", "root", "账号")
	s.cmd.Flags().StringVarP(&s.Password, "password", "", "123456", "密码")
	s.cmd.Flags().StringVarP(&s.Dbname, "name", "", "blog", "数据库名称")
	s.cmd.Flags().StringVarP(&s.Config, "config", "", "charset=utf8mb4&parseTime=True&loc=Local", "数据库配置")
}

func (s *MigrateCmd) RunMigrate(cmd *cobra.Command, args []string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", s.Username, s.Password, s.Host, s.Port, "", s.Config)
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
	err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", s.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	// 创建数据库（如果不存在）
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", s.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(fmt.Sprintf("USE `%s`", s.Dbname)).Error
	if err != nil {
		log.Fatal(err)
	}

	switch s.action {
	case "migrate":
		s.MigrateDatabase(db)
	case "reset":
		s.ResetDatabase(db)
	default:
		log.Fatal("action not support")
	}
}

// 迁移数据库
func (s *MigrateCmd) MigrateDatabase(db *gorm.DB) {
	// 读取 SQL 文件内容
	content, err := os.ReadFile(s.sqlFile)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	log.Println("Database initialized successfully.")
}

// 重置数据库
func (s *MigrateCmd) ResetDatabase(db *gorm.DB) {
	var err error
	// ****重置登录历史表****
	err = ClearTable(db, entity.TableNameUserLoginHistory)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置角色菜单表****
	err = ClearTable(db, entity.TableNameRoleMenu)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置角色接口表****
	err = ClearTable(db, entity.TableNameRoleApi)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置菜单表****
	err = ClearTable(db, entity.TableNameMenu)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置接口表****
	err = ClearTable(db, entity.TableNameApi)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置操作记录表****
	err = ClearTable(db, entity.TableNameOperationLog)
	if err != nil {
		log.Fatal(err)
	}

	// ****重置上传记录表****
	err = ClearTable(db, entity.TableNameUploadRecord)
	if err != nil {
		log.Fatal(err)
	}
}

func ClearTable(db *gorm.DB, tableName string) (err error) {
	// 清空表的数据
	err = db.Exec(fmt.Sprintf("DELETE FROM `%v`", tableName)).Error
	if err != nil {
		return err
	}
	// 重置 AUTO_INCREMENT 值为 1
	err = db.Exec(fmt.Sprintf("ALTER TABLE `%v` AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		return err
	}
	return nil
}
