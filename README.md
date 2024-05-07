# Install

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

# Generate

```
protoc --go_out=. --go-grpc_out=. ./coinmeca_grpc.proto
```

# Error Fix

```
error-protoc-gen-go-program-not-found-or-is-not-executable
```

https://stackoverflow.com/questions/57700860/
