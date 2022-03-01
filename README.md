# gRPC with Go
Unary

## Run Server :

go run main.go


## Run Client :

go run main.go --name=John


## Generate from protocol buffers with protoc :

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/chat.proto

