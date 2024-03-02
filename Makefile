.PHONY: build-protos clean-protos build test fmt vet lint tidy clean-bazel clean fix-build leetcode

build-protos:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/proto/example/example.proto

clean-protos:
	rm internal/proto/example/*.pb.go

build: build-protos
	bazel build //...

test:
	bazel test --test_output=errors //...

fmt:
	gofmt -s -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy -v

# See https://github.com/bazelbuild/bazel-gazelle
update-bazel: build-protos
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
	bazel run //:gazelle

clean-bazel:
	bazel clean --expunge

clean: clean-bazel clean-protos

fix-build: fmt tidy update-bazel

leetcode:
	chmod +x ./scripts/leetcode.sh
	./scripts/leetcode.sh