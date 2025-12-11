.PHONY: help deps gen-gin gen-model gen-model-dsn gen-ts gen-all clean

# 默认目标
help:
	@echo "可用命令："
	@echo "  make deps          - 安装依赖"
	@echo "  make gen-gin       - 生成 Gin 框架代码"
	@echo "  make gen-model     - 从 SQL 文件生成数据库模型代码"
	@echo "  make gen-model-dsn - 从数据库生成模型代码（需配置 DSN）"
	@echo "  make gen-ts        - 生成 TypeScript 代码"
	@echo "  make gen-all       - 生成所有代码"
	@echo "  make clean         - 清理生成的代码"

# 安装依赖
deps:
	go mod tidy
	go mod download

# 生成 Gin 框架代码（博客前台）
gen-gin-blog:
	@echo "生成博客前台 types..."
	go run main.go api types -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -o='../blog-gin/api/blog' -n='%s.go'
	@echo "生成博客前台 logic..."
	go run main.go api logic -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.lg.go'
	@echo "生成博客前台 handler..."
	go run main.go api handler -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.hdl.go'
	@echo "生成博客前台 router..."
	go run main.go api router -f='../blog-gozero/service/api/blog/proto/blog.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/blog' -n='%s.rt.go'
	@echo "博客前台代码生成完成！"

# 生成 Gin 框架代码（管理后台）
gen-gin-admin:
	@echo "生成管理后台 types..."
	go run main.go api types -f='../blog-gozero/service/api/admin/proto/admin.api' -t='./template/gin' -o='../blog-gin/api/admin' -n='%s.go'
	@echo "生成管理后台 logic..."
	go run main.go api logic -f='../blog-gozero/service/api/admin/proto/admin.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/admin' -n='%s.lg.go'
	@echo "生成管理后台 handler..."
	go run main.go api handler -f='../blog-gozero/service/api/admin/proto/admin.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/admin' -n='%s.hdl.go'
	@echo "生成管理后台 router..."
	go run main.go api router -f='../blog-gozero/service/api/admin/proto/admin.api' -t='./template/gin' -c='github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx' -o='../blog-gin/api/admin' -n='%s.rt.go'
	@echo "管理后台代码生成完成！"

# 生成所有 Gin 框架代码
gen-gin: gen-gin-blog gen-gin-admin
	@echo "所有 Gin 框架代码生成完成！"

# 从 SQL 文件生成数据库模型代码
gen-model:
	@echo "从 SQL 文件生成数据库模型代码..."
	go run main.go model ddl -t=./template/go-zero/model.tpl -n='%v_model.go' -o='./runtime/model' -s='./testdata/test.sql'
	@echo "数据库模型代码生成完成！"

# 从数据库生成模型代码（需要配置 DSN）
gen-model-dsn:
	@echo "从数据库生成模型代码..."
	@echo "请确保已配置正确的数据库连接信息"
	go run main.go model dsn -t=./template/go-zero/model.tpl -n='%v_model.go' -o='./runtime/model' -s='root:password@(localhost:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local'
	@echo "数据库模型代码生成完成！"

# 从数据库生成模型代码到 blog-gozero 项目（生产环境）
gen-model-prod:
	@echo "从生产数据库生成模型代码..."
	@read -p "请输入数据库连接字符串 (DSN): " dsn; \
	go run main.go model dsn -t=./template/go-zero/model.tpl -n='%v_model.go' -o='../blog-gozero/service/model' -s="root:mysql7914@(veweiyi.cn:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local"
	@echo "数据库模型代码生成完成！"

# 生成 TypeScript 代码（博客前台）
gen-ts-blog:
	@echo "生成博客前台 TypeScript 代码..."
	go run main.go web typescript -n='%v.ts' -t='./template/web' -o='./runtime/blog/api' -m='api' -i='IApiResponse' -f='../blog-gozero/service/api/blog/proto/blog.api'
	@echo "博客前台 TypeScript 代码生成完成！"

# 生成 TypeScript 代码（管理后台）
gen-ts-admin:
	@echo "生成管理后台 TypeScript 代码..."
	go run main.go web typescript -n='%v.ts' -t='./template/web' -o='./runtime/admin/api' -m='api' -i='IApiResponse' -f='../blog-gozero/service/api/admin/proto/admin.api'
	@echo "管理后台 TypeScript 代码生成完成！"

# 生成所有 TypeScript 代码
gen-ts: gen-ts-blog gen-ts-admin
	@echo "所有 TypeScript 代码生成完成！"

# 生成所有代码（不包括从数据库生成）
gen-all: gen-gin gen-model gen-ts
	@echo "所有代码生成完成！"

# 清理生成的代码
clean:
	@echo "清理生成的代码..."
	rm -rf ./runtime/model
	rm -rf ./runtime/blog
	rm -rf ./runtime/admin
	@echo "清理完成！"
