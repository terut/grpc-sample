### Requirement

- Protocol Buffers (See: https://developers.google.com/protocol-buffers/)

```sh
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

### Preparation

```
$ protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:proto user.proto
$ protoc -I./proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:proto user.proto
```

