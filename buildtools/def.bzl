load("@rules_proto//proto:defs.bzl", "proto_library")
load("//buildtools/internal/proto-options:defs.bzl", "proto_options")

def proto_library(
        name,
        srcs,
        deps = [],
        visibility = [],
        plugins = []):
    proto_library(
        name = name,
        srcs = srcs,
        deps = deps,
        visibility = visibility,
        plugins = plugins,
    )
