"""Module extension for downloading prebuilt tool binaries."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_file")

# shfmt releases: https://github.com/mvdan/sh/releases
SHFMT_VERSION = "3.12.0"
SHFMT_BINARIES = {
    "darwin_amd64": "c31548693de6584e6164b7ed5fbb7b4a083f2d937ca94b4e0ddf59aa461a85e4",
    "darwin_arm64": "d903802e0ce3ecbc82b98512f55ba370b0d37a93f3f78de394f5b657052b33dd",
    "linux_amd64": "d9fbb2a9c33d13f47e7618cf362a914d029d02a6df124064fff04fd688a745ea",
    "linux_arm64": "5f3fe3fa6a9f766e6a182ba79a94bef8afedafc57db0b1ad32b0f67fae971ba4",
}

# dprint releases: https://github.com/dprint/dprint/releases
DPRINT_VERSION = "0.50.2"
DPRINT_BINARIES = {
    "darwin_amd64": "61becbf8d1b16540e364a4f00be704266ae322ee0ff3ba66a4a21033f66a8d55",
    "darwin_arm64": "f534bcc054947ab2a42c069b5f6027914d252729bd15c1109812313b35a662a5",
    "linux_amd64": "95c7e633a67531ffc4990c152d59ed0802e1c0caf7e27e424e9cea9ef3d499d4",
    "linux_arm64": "039d4dca4360cb6622a2b56c3fc29ea71c356cd954e0b9566bff1a70e75beda8",
}

def _shfmt_repo_impl(ctx):
    """Repository rule that creates a shfmt binary for the current platform."""
    os = ctx.os.name
    arch = ctx.os.arch

    if os == "mac os x":
        os = "darwin"
    elif os.startswith("windows"):
        os = "windows"
    else:
        os = "linux"

    if arch == "aarch64" or arch == "arm64":
        arch = "arm64"
    else:
        arch = "amd64"

    platform = "{}_{}".format(os, arch)
    sha256 = SHFMT_BINARIES.get(platform)

    if not sha256:
        fail("Unsupported platform: {}".format(platform))

    url = "https://github.com/mvdan/sh/releases/download/v{version}/shfmt_v{version}_{platform}".format(
        version = SHFMT_VERSION,
        platform = platform,
    )

    ctx.download(
        url = url,
        output = "shfmt",
        sha256 = sha256,
        executable = True,
    )

    ctx.file("BUILD.bazel", """
exports_files(["shfmt"])

sh_binary(
    name = "shfmt_bin",
    srcs = ["shfmt"],
    visibility = ["//visibility:public"],
)
""")

shfmt_repo = repository_rule(
    implementation = _shfmt_repo_impl,
    attrs = {},
)

def _dprint_repo_impl(ctx):
    """Repository rule that creates a dprint binary for the current platform."""
    os = ctx.os.name
    arch = ctx.os.arch

    if os == "mac os x":
        os = "darwin"
        os_dprint = "apple-darwin"
    elif os.startswith("windows"):
        fail("dprint: Windows not supported")
    else:
        os = "linux"
        os_dprint = "unknown-linux-gnu"

    if arch == "aarch64" or arch == "arm64":
        arch = "arm64"
        arch_dprint = "aarch64"
    else:
        arch = "amd64"
        arch_dprint = "x86_64"

    platform = "{}_{}".format(os, arch)
    sha256 = DPRINT_BINARIES.get(platform)

    if not sha256:
        fail("Unsupported platform: {}".format(platform))

    url = "https://github.com/dprint/dprint/releases/download/{version}/dprint-{arch}-{os}.zip".format(
        version = DPRINT_VERSION,
        arch = arch_dprint,
        os = os_dprint,
    )

    ctx.download_and_extract(
        url = url,
        sha256 = sha256,
    )

    ctx.file("BUILD.bazel", """
exports_files(["dprint"])

sh_binary(
    name = "dprint_bin",
    srcs = ["dprint"],
    visibility = ["//visibility:public"],
)
""")

dprint_repo = repository_rule(
    implementation = _dprint_repo_impl,
    attrs = {},
)

def _prebuilt_tools_impl(ctx):
    shfmt_repo(name = "shfmt_prebuilt")
    dprint_repo(name = "dprint_prebuilt")

prebuilt_tools = module_extension(
    implementation = _prebuilt_tools_impl,
)
