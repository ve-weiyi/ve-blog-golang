
# 生成service.go
go run main.go api gin \
-n='%v.ts' \
-t='./resource/template/gin' \
-o='./runtime/blog/api'  \
-m='api' \
-i='IApiResponse' \
-f='/Users/weiyi/Github/ve-blog-golang/zero/service/api/blog/proto/blog.api'
