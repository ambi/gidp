clean:
	rm -f echoapp grpcapp

echoapp:
	go build -o echoapp infra/cmd/echoapp/main.go

grpcapp:
	go build -o grpcapp infra/cmd/grpcapp/main.go

protoc:
  # go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc -I=adapter/rpccontroller --go_out=adapter/rpccontroller adapter/rpccontroller/api.proto
	protoc -I=adapter/rpccontroller --go-grpc_out=adapter/rpccontroller adapter/rpccontroller/api.proto

test:
	go test ./...
