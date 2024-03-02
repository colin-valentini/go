# internal/proto

A clear, simple Go-based proto gRPC example.

## prerequisites

Install core proto compiler:

```sh
brew install protoc
```

Install Go proto compiler and gRPC plugins:

```sh
brew install protoc-gen-go protoc-gen-go-grpc
```

## compiling the protos

First, **change into the `example/` subdirectory**, then run:

```sh
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  example.proto
```

Since there are external (Google) libraries that are needed for the generated
Go code, you should additionally run `go get .` to fetch the dependencies and
update the `go.mod` file at the root of this respository.

The protos can then be imported and used like so:

```go
import pb "github.com/colin-valentini/go/internal/proto/example"

var req = &pb.GreetingRequest{Id: 123}
var res = &pb.GreetingResponse{Message: "Hello world!"}
```

### bazel

See the following links about compatability with Bazel and Gazelle:

- [Proto Depedencies](https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#proto-dependencies)
- [gRPC Dependencies](https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#grpc-dependencies)

## demo

First, **change into the `example/` subdirectory**, start the server:

```sh
go run server/main.go
```

In a separate terminal session, run the client side main:

```sh
go run client/main.go
```

### bazel

Start the server:

```sh
bazel run //internal/proto/example/server:main
```

In a separate terminal session, run the client side main:

```sh
bazel run //internal/proto/example/client:main
```
