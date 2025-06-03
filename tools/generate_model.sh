#-f 选项用于强制删除文件而不提示确认，-r 选项用于递归删除目录中的文件。
#rm -f '/Users/weiyi/Github/ve-blog-golang/blog-gozero/service/model'
#goctl model mysql ddl -src  test.sql -dir ../ --style go_zero -c --home ../../../../template


# dsn(Data Source Name)线上数据库生成model文件
#go run main.go model dsn \
#-t=./template/go-zero/model.tpl \
#-n='%v_model.go' \
#-o='/Users/weiyi/Github/ve-blog-golang/blog-gozero/service/model'  \
#-s='root:mysql7914@(veweiyi.cn:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local'


# dsn(Data Source Name)本地数据库生成model文件
go run main.go model dsn \
-t=./template/go-zero/model.tpl \
-n='%v_model.go' \
-o='/Users/weiyi/Github/ve-blog-golang/blog-gozero/service/model'  \
-s='root:mysql7914@(veweiyi.cn:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local'

# ddl(Data Definition Language)生成model文件
#go run main.go model ddl \
#-t=./template/go-zero/model.tpl \
#-n='%v_model.go' \
#-o='./runtime/model'  \
#-s='./testdata/test.sql'
