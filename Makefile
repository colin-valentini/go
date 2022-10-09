.PHONY: build test fmt vet lint tidy fix-build

build:
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

update-bazel:
	bazel run //:gazelle-update-repos
	bazel run //:gazelle

clean-bazel:
	bazel clean --expunge

fix-build: fmt tidy update-bazel
