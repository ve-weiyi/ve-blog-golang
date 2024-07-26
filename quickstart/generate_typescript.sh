
# 生成api.ts
go run main.go web typescript \
-n='%v.ts' \
-t='./resource/template/web' \
-o='./runtime/blog/api'  \
-m='api' \
-i='IApiResponse' \
-f='/Users/weiyi/Github/ve-blog-golang/zero/service/api/blog/proto/blog.api'
