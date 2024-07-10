#-f 选项用于强制删除文件而不提示确认，-r 选项用于递归删除目录中的文件。
#rm -f '/Users/weiyi/Github/ve-blog-golang/zero/service/model'
#goctl model mysql ddl -src  test.sql -dir ../ --style go_zero -c --home ../../../../resource/template


# ddl生成entity文件
#go run main.go model ddl \
#-t=./temp/model.tpl \
#-n='%s.go' \
#-o=./tmp \
#-s=./test.sql

# dsn(data source name)生成entity文件
go run main.go model dsn \
-t=./resource/template/go-zero/model.tpl \
-n='%v_model_gen.go' \
-o='/Users/weiyi/Github/ve-blog-golang/zero/service/model'  \
-s='root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local'

