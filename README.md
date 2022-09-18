# go

![build](https://github.com/colin-valentini/go/actions/workflows/build.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/colin-valentini/go/branch/main/graph/badge.svg?token=EZ72S8AIU9)](https://codecov.io/gh/colin-valentini/go)

This repository is a collection of Go practice problems, or other assorted findings.

## prequisites

Requires [golang](https://formulae.brew.sh/formula/go#default),
[bazel](https://bazel.build/install/bazelisk), and
[golangci-lint](https://golangci-lint.run/usage/install/#macos).

## local development

Run `make build` to build all of the things, `make test` to test all of the
things, and `make lint` to lint all of the things. If you get a bazel build
error, try running `make fix-build` which will run formatting and build file
updates. 

## build tooling

This repository uses [Bazel](https://bazel.build/) to compile and run tests
locally. Bazel can be downloaded, installed, and managed on your local machine
using [Bazelisk](https://github.com/bazelbuild/bazelisk).

Bazel `BUILD` files are automatically generated and managed using
[Gazelle](https://github.com/bazelbuild/bazel-gazelle).

## linter

Linting is handled using [golangci-lint](https://golangci-lint.run/) and is
configured using the [`.golangci.yml`](.golangci.yml) file.

## ci/cd pipeline

The CI/CD pipeline is run via GitHub actions and does not use Bazel, rather the
standard Golang toolchain to build, vet, lint, and test. Test coverage is
uploaded to Codecov for reports on coverage deltas.

See:
- [actions/setup-go](actions/setup-go)
- [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action)
- [codecov/codecov-action](https://github.com/codecov/codecov-action)

