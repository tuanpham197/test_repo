## Generate proto

### Document quick start
[Documentation](https://grpc.io/docs/languages/go/quickstart/)

### Run generate proto for comment
```
protoc --proto_path=./ --go_out=./comment --go_opt=paths=source_relative \
    --go-grpc_out=./comment --go-grpc_opt=paths=source_relative \
    comment.proto
```

### Add tag version for pkg
1. Add tag
```
git tag v0.1.0
```
2. Push
```
git push origin v0.1.0
```
3. Go git update
```
go get -u github.com/tuanpham197/test_repo@v0.1.0
```