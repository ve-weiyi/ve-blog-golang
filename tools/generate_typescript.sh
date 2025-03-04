
# 生成api.ts
go run main.go web typescript \
-n='%v.ts' \
-t='./template/web' \
-o='./runtime/blog/api'  \
-m='api' \
-i='IApiResponse' \
-f='/Users/weiyi/Github/ve-blog-golang/blog-gozero/service/api/blog/proto/blog.api'

# 生成api.ts
go run main.go web typescript \
-n='%v.ts' \
-t='./template/web' \
-o='./runtime/admin/api'  \
-m='api' \
-i='IApiResponse' \
-f='/Users/weiyi/Github/ve-blog-golang/blog-gozero/service/api/admin/proto/admin.api'
