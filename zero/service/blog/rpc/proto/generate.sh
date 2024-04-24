rm ../blog.go
goctl rpc protoc ./blog.proto --go_out=../pb --go-grpc_out=../pb --zrpc_out=../ -m --style go_zero
