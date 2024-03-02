.PHONY: build test fmt vet lint tidy fix-build leetcode

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

# See https://github.com/bazelbuild/bazel-gazelle
update-bazel:
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
	bazel run //:gazelle

clean-bazel:
	bazel clean --expunge

fix-build: fmt tidy update-bazel

leetcode:
	chmod +x ./scripts/leetcode.sh
	./scripts/leetcode.sh