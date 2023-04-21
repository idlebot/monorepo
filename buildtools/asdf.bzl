""" asdf.bzl contains tool_versions repository rule.

tool_versions rule generates .tool-versions file in the workspace directory
and can be used to keep Bazel and asdf tool versions in sync.

Example:

Add the following to your WORKSPACE file:

GO_VERSION = "1.19.6"
RULES_GO_VERSION = "0.38.1"

load("//buildtools:asdf.bzl", "tool_versions")

tool_versions(
    name = "asdf",
    versions = {
        "golang": GO_VERSION,
    },
    workspace_dir = __workspace_dir__,
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(RULES_GO_VERSION),
        "https://github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(RULES_GO_VERSION),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version = GO_VERSION)

And then run

bazel run @asdf//:gen_tool_versions

to generate a .tool-versions file in the workspace directory.

with the following contents:

golang 1.19.6

Note that since we use the GO_VERSION variable in both tool_versions
and go_register_toolchain, we ensure that both Bazel and asdf will use
the same version for this repository.

 """

def _tool_versions_impl(ctx):
    """ _tool_versions_impl is the implementation of tool_versions rule. """
    content = []

    for tool in ctx.attr.versions.keys():
        content.append("{0} {1}".format(tool, ctx.attr.versions[tool]))

    ctx.file(".tool-versions", "\n".join(content))

    content = [
        "#!/usr/bin/env bash",
        "cp {0} .".format(ctx.path(".tool-versions")),
    ]

    ctx.file(
        "tool_versions",
        content = "\n".join(content),
        executable = True,
    )
    ctx.file("BUILD.bazel", 'exports_files(["tool_versions"])')

    workspace_dir = ctx.workspace_root.basename
    return ctx.execute(
        [
            ctx.path("tool_versions"),
        ],
        working_directory = workspace_dir,
    )

_tool_versions = repository_rule(
    implementation = _tool_versions_impl,
    attrs = {
        "versions": attr.string_dict(
            allow_empty = False,
        ),
    },
)

def gen_tool_versions(name, versions):
    """ gen_tool_versions generates .tool-versions file in the workspace directory. """
    _tool_versions(
        name = name,
        versions = versions,
    )
