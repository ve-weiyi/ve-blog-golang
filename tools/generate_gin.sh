
# api生成types文件
go run main.go api types \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-o='../blog-gin/api/blog'  \
-n='%s.go'

# api生成logic文件
go run main.go api logic \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/blog'  \
-n='%s.lg.go'

# api生成handler文件
go run main.go api handler \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/blog'  \
-n='%s.hdl.go'

# api生成router文件
go run main.go api router \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/blog'  \
-n='%s.rt.go'


###############################


# api生成types文件
go run main.go api types \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-o='../blog-gin/api/admin'  \
-n='%s.go'

# api生成logic文件
go run main.go api logic \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/admin'  \
-n='%s.lg.go'

# api生成handler文件
go run main.go api handler \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/admin'  \
-n='%s.hdl.go'

# api生成router文件
go run main.go api router \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/api/admin'  \
-n='%s.rt.go'
