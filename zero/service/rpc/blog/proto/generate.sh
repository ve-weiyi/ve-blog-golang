#
#goctl rpc protoc ./account.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/account.yaml
#rm ../account.go
#
#goctl rpc protoc ./permission.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/permission.yaml
#rm ../permission.go
#
#goctl rpc protoc ./article.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/article.yaml
#rm ../article.go

goctl rpc protoc ./comment.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
rm -f ../etc/comment.yaml
rm ../comment.go
#
#goctl rpc protoc ./remark.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/remark.yaml
#rm ../remark.go
#
#goctl rpc protoc ./friend.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/friend.yaml
#rm ../friend.go
#
#goctl rpc protoc ./photo.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/photo.yaml
#rm ../photo.go
#
#goctl rpc protoc ./syslog.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/syslog.yaml
#rm ../syslog.go
#
#goctl rpc protoc ./talk.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/talk.yaml
#rm ../talk.go
#
#goctl rpc protoc ./website.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/website.yaml
#rm ../website.go
#
#goctl rpc protoc ./config.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/config.yaml
#rm ../config.go
#
#
#goctl rpc protoc ./chat.proto --go_out=../internal/pb --go-grpc_out=../internal/pb --zrpc_out=../ -m --style go_zero
#rm -f ../etc/chat.yaml
#rm ../chat.go
