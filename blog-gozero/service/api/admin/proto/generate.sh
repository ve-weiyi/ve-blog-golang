# 根据 api 文件生成 业务代码
goctl api go -api admin.api -dir ../ --style go_zero --home ../../../../.goctl/template

# 根据 api 文件生成 swagger 文档。
goctl api swagger --api admin.api --dir ../docs
