apt update
apt install -y protobuf-compiler
cd protoc
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./recipe.proto