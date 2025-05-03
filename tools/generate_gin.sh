
# api生成dto文件
go run main.go api dto \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-o='../blog-gin/service/blog'  \
-n='%s.go'

# api生成service文件
go run main.go api service \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/blog'  \
-n='%s.sv.go'

# api生成controller文件
go run main.go api controller \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/blog'  \
-n='%s.ctl.go'

# api生成router文件
go run main.go api router \
-f='../blog-gozero/service/api/blog/proto/blog.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/blog'  \
-n='%s.rt.go'


###############################

# api生成dto文件
go run main.go api dto \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-o='../blog-gin/service/admin'  \
-n='%s.go'

# api生成service文件
go run main.go api service \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/admin'  \
-n='%s.sv.go'

# api生成controller文件
go run main.go api controller \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/admin'  \
-n='%s.ctl.go'

# api生成router文件
go run main.go api router \
-f='../blog-gozero/service/api/admin/proto/admin.api' \
-t='./template/gin' \
-c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx'  \
-o='../blog-gin/service/admin'  \
-n='%s.rt.go'
