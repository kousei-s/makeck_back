apt update
apt install -y protobuf-compiler
cd protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./recipe.proto