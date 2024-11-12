build-server:
	go run gin/main.go

build-blog-rpc:
	go run gozero/service/rpc/blog/blog.go

build-blog-api:
	go run gozero/service/api/blog/blog.go

build-admin-api:
	go run gozero/service/api/admin/admin.go