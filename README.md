## Generate proto

### Document quick start
[Documentation](https://grpc.io/docs/languages/go/quickstart/)

### Run generate proto for comment
```
protoc --proto_path=./ --go_out=./comment --go_opt=paths=source_relative \
    --go-grpc_out=./comment --go-grpc_opt=paths=source_relative \
    comment.proto
```