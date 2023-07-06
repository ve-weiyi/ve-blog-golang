package dbdriver

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

// GEN 自动生成 GORM 模型结构体文件及使用示例 https://blog.csdn.net/Jeffid/article/details/126898000
const dsn = "root:mysql7914@(127.0.0.1:3306)/blog-plus?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "tb_",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	log.Println("mysql connection done")
}

func TestDBService(t *testing.T) {
	var driver DBDriver
	driver = &MysqlDriver{DB: db}
	//dbList, err := driver.GetDB()
	//if err != nil {
	//	log.Println("-->", err)
	//	return
	//}
	//log.Println("dbList-->", jsonconv.ObjectToJsonIndent(dbList))
	//
	//tabelList, err := driver.GetTables("blog-v2")
	//if err != nil {
	//	log.Println("-->", err)
	//	return
	//}
	//log.Println("tabelList-->", jsonconv.ObjectToJsonIndent(tabelList))

	columnList, err := driver.GetTableColumns("blog-v2", "user_account")
	if err != nil {
		log.Println("-->", err)
		return
	}
	log.Println("columnList-->", db.Migrator().CurrentDatabase(), jsonconv.ObjectToJsonIndent(columnList))
}
