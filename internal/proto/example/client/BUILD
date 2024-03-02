load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "client_lib",
    srcs = ["main.go"],
    importpath = "github.com/colin-valentini/go/internal/proto/example/client",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/proto/example",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
