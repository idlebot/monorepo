"""
csharp_gapic_library rule
"""

load("@rules_gapic//:gapic.bzl", "proto_custom_library")
load("@rules_gapic//csharp:csharp_gapic.bzl", _csharp_grpc_library = "csharp_grpc_library", _csharp_proto_library = "csharp_proto_library")
load("@rules_proto//proto:defs.bzl", "ProtoInfo")

def csharp_proto_library(name, deps, **kwargs):
    _csharp_proto_library(name, deps, **kwargs)

def csharp_grpc_library(name, srcs, deps, **kwargs):
    _csharp_grpc_library(name, srcs, deps, **kwargs)

def csharp_gapic_library(
        name,
        srcs,
        deps = [],
        namespace = None,
        visibility = None,
        grpc_service_config = None,
        common_resources_config = None,
        service_yaml = None,
        rest_numeric_enums = False,
        transport = None,
        generator_binary = "@gapic_generator_csharp//rules_csharp_gapic:csharp_gapic_generator_binary",
        target_frameworks = ["net7.0"],
        **kwargs):
    """Generates C# client library from proto files.

    Args:
        name: name of the target
        srcs: list of proto_library targets
        deps: list of dependencies
        namespace: namespace of the generated library
        visibility: list of visibility rules
        grpc_service_config: path to gRPC service config
        common_resources_config: path to common resources config
        service_yaml: path to service yaml
        rest_numeric_enums: whether to use numeric values for enums in REST
        transport: transport to use, either "grpc" or "rest"
        generator_binary: path to the generator binary
        target_frameworks: list of target frameworks to build for
        **kwargs: additional arguments
    """

    # Build zip file of gapic-generator output
    srcjar_name = "{name}_srcjar".format(name = name)
    raw_srcjar_name = srcjar_name + "_raw"
    name_gapicinfo = "{name}_gapicinfo".format(name = name)
    output_suffix = ".srcjar"
    opt_file_args = {}
    opt_args = []
    if grpc_service_config:
        opt_file_args[grpc_service_config] = "grpc-service-config"
    if common_resources_config:
        opt_file_args[common_resources_config] = "common-resources-config"
    if service_yaml:
        opt_file_args[service_yaml] = "service-config"
    if rest_numeric_enums:
        opt_args.append("rest-numeric-enums={}".format(rest_numeric_enums))
    if transport:
        opt_args.append("transport={}".format(transport))

    proto_custom_library(
        name = raw_srcjar_name,
        deps = srcs,
        plugin = Label(generator_binary),
        visibility = visibility,
        opt_file_args = opt_file_args,
        opt_args = opt_args,
        output_type = "gapic",
        output_suffix = output_suffix,
        **kwargs
    )

    sources_name = "{name}_sources".format(name = name)
    main_file = ":%s" % srcjar_name + output_suffix
    main_dir = "%s_main" % srcjar_name
    _csharp_gapic_sources(
        name = sources_name,
        srcjar = main_file,
    )

    print(main_file)
    print(main_dir)

    # csharp_srcs_name = name + "_csharp_srcs"
    # _csharp_gapic_sources(
    #     name = csharp_srcs_name,
    #     srcjar = ":%s.cs" % main_file,
    # )

    # csharp_library(
    #     name = name,
    #     srcs = [":%s" % csharp_srcs_name],
    #     private_deps = [
    #         "@paket.main//google.api.commonprotos",
    #         "@paket.main//google.api.gax",
    #         "@paket.main//google.api.gax.grpc",
    #         "@paket.main//google.protobuf",
    #         "@paket.main//grpc.core",
    #         "@paket.main//grpc.core.api",
    #         "@paket.main//microsoft.extensions.dependencyinjection.abstractions",
    #         "@paket.main//microsoft.extensions.logging.abstractions",
    #         "@paket.main//microsoft.netcore.app.ref",
    #     ],
    #     target_frameworks = target_frameworks,
    #     visibility = visibility,
    #     deps = deps,
    # )

def _csharp_gapic_sources_impl(ctx):
    gapic_srcjar = ctx.file.srcjar

    srcs = ctx.files.srcjar
    csharp_srcs = []
    print(srcs)
    print(ctx.attr)
    print(ctx.build_file_path)
    print(ctx.outputs)
    print(ctx.files)
    print(ctx.genfiles_dir)
    print(ctx.attr.srcjar)

    script = """
    unzip -q {gapic_srcjar} -d {output_dir_path}
    """.format(
        gapic_srcjar = gapic_srcjar.path,
        output_dir_path = "hello",
    )

    return [DefaultInfo(files = depset(csharp_srcs))]

_csharp_gapic_sources = rule(
    _csharp_gapic_sources_impl,
    attrs = {
        "srcjar": attr.label(allow_files = True),
    },
)

# def _go_gapic_postprocessed_srcjar_impl(ctx):
#     go_ctx = go_context(ctx)

#     gapic_srcjar = ctx.file.gapic_srcjar
#     output_main = ctx.outputs.main
#     output_test = ctx.outputs.test
#     output_metadata = ctx.outputs.metadata

#     output_dir_name = ctx.label.name
#     output_dir_path = "%s/%s" % (output_main.dirname, output_dir_name)

#     formatter = _get_gofmt(go_ctx)

#     script = """
#     unzip -q {gapic_srcjar} -d {output_dir_path}
#     {formatter} -w -l {output_dir_path}
#     pushd .
#     cd {output_dir_path}
#     zip -q -r {output_dir_name}-test.srcjar . -i ./*_test.go
#     find . -name "*_test.go" -delete
#     zip -q -r {output_dir_name}.srcjar . -i ./*.go
#     find . -name "*.go" -delete
#     zip -q -r {output_dir_name}-metadata.srcjar . -i ./*.json
#     popd
#     mv {output_dir_path}/{output_dir_name}-test.srcjar {output_test}
#     mv {output_dir_path}/{output_dir_name}.srcjar {output_main}
#     mv {output_dir_path}/{output_dir_name}-metadata.srcjar {output_metadata}
#     """.format(
#         gapic_srcjar = gapic_srcjar.path,
#         output_dir_name = output_dir_name,
#         output_dir_path = output_dir_path,
#         formatter = formatter.path,
#         output_main = output_main.path,
#         output_test = output_test.path,
#         output_metadata = output_metadata.path,
#     )

#     ctx.actions.run_shell(
#         inputs = [gapic_srcjar],
#         tools = [formatter],
#         command = script,
#         outputs = [output_main, output_test, output_metadata],
#     )

# _go_gapic_postprocessed_srcjar = rule(
#     _go_gapic_postprocessed_srcjar_impl,
#     attrs = {
#         "gapic_srcjar": attr.label(mandatory = True, allow_single_file = True),
#         "_go_context_data": attr.label(
#             default = "@io_bazel_rules_go//:go_context_data",
#         ),
#     },
#     outputs = {
#         "main": "%{name}.srcjar",
#         "test": "%{name}-test.srcjar",
#         "metadata": "%{name}-metadata.srcjar",
#     },
#     toolchains = ["@io_bazel_rules_go//go:toolchain"],
# )

def _debug(ctx):
    print(ctx.attr.src[ProtoInfo].direct_descriptor_set.short_path)
    return [DefaultInfo()]

debug = rule(
    implementation = _debug,
    attrs = {
        "src": attr.label(
            mandatory = True,
            providers = [ProtoInfo],
        ),
    },
)
