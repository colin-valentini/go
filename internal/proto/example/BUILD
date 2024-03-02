load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "example_proto",
    srcs = ["example.proto"],
    visibility = ["//:__subpackages__"],
)

go_proto_library(
    name = "example_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/colin-valentini/go/internal/proto/example",
    proto = ":example_proto",
    visibility = ["//:__subpackages__"],
)

go_library(
    name = "example",
    embed = [":example_go_proto"],
    importpath = "github.com/colin-valentini/go/internal/proto/example",
    visibility = ["//:__subpackages__"],
)
