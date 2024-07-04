
# ddl生成entity文件
#go run main.go model ddl \
#-t=./temp/model.tpl \
#-n='%s.go' \
#-o=./tmp \
#-s=./test.sql

# dsn(data source name)生成entity文件
#go run main.go model dsn \
#-t=./temp/model.tpl \
#-n='%s.go' \
#-o=/Users/weiyi/Github/ve-blog-golang/server/api/blog/model/entity  \
#-s='root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local'

# api生成router文件
go run main.go api router \
-t=./temp \
-n='%s.go' \
-o='/Users/weiyi/Github/ve-blog-golang/server/api/blog'  \
-f='/Users/weiyi/Github/ve-blog-golang/zero/service/blog/api/proto/blog.api'
