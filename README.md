# Install

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

# Generate

```
protoc --go_out=. --go-grpc_out=. ./coinmeca_grpc.proto
```