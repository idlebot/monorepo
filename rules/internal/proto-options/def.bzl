"""
proto_options rule
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
        "src": attr.label(
            mandatory = True,
            providers = [ProtoInfo],
        ),
    },
)
