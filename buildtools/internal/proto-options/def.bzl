"""
proto_options rule uses the proto-options utility to automate the generation of
standard lanaguage options in .proto files

It also returns a ProtoOptionsInfo provider that contains the generated option
values that can be used by other rules, in particular the gapic rules.
"""

load("@rules_gapic//:gapic.bzl", "proto_custom_library")
load("@rules_gapic//csharp:csharp_gapic.bzl", _csharp_grpc_library = "csharp_grpc_library", _csharp_proto_library = "csharp_proto_library")
load("@rules_proto//proto:defs.bzl", "ProtoInfo")

def _proto_options_impl(ctx):
    print(ctx.attr.src[ProtoInfo].direct_descriptor_set.short_path)
    return [DefaultInfo()]

proto_options = rule(
    implementation = _proto_options_impl,
    attrs = {
        "srcs": attr.label_list(allow_files = [".proto"]),
        "deps": attr.label_list(providers = [ProtoInfo]),
        "_compiler": attr.label(
            default = Label("//tools:example_compiler"),
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
    },
)
