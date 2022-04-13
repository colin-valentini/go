# go

![build](https://github.com/colin-valentini/go/actions/workflows/build.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/colin-valentini/go/branch/main/graph/badge.svg?token=EZ72S8AIU9)](https://codecov.io/gh/colin-valentini/go)

This repository is a collection of Go practice problems, or other assorted findings.

## prequisites

Requires [golang](https://formulae.brew.sh/formula/go#default), [bazel](https://bazel.build/install/bazelisk), and [golangci-lint](https://golangci-lint.run/usage/install/#macos).

## local development

Run `make build` to build all of the things, and `make test` to test all of the things. If you get a bazel build error, try running `make fix-build` which will run formatting and build file updates. If you get a lint error, run `make lint` locally to execute the linter — see `.golangci.yml` for configuration, and https://golangci-lint.run/ for docs.
