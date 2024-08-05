# 根据 api 文件生成 业务代码
goctl api go -api blogrpc.api -dir ../ --style go_zero --home ../../../../resource/template

# 根据 api 文件生成 swagger 文档。
#goctl api plugin -plugin goctl-swagger="swagger -filename blogrpc.json" -api blogrpc.api -dir ./
