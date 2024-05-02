# Install

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

# Generate

```
protoc --go_out=. --go-grpc_out=. ./coinmeca_grpc.proto
```

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable
