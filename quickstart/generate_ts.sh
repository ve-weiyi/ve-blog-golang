
# 生成api.ts
go run main.go web typescript \
-n='%v.ts' \
-t='./resource/template/web' \
-o='./runtime/api'  \
-m='api' \
-f='/Users/weiyi/Github/ve-blog-golang/zero/service/api/blog/proto/admin.api'

