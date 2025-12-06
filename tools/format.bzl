"""Macro for creating a multi-target format runner."""

load("@bazel_skylib//rules:write_file.bzl", "write_file")

def format_multi(name, targets, **kwargs):
    """Creates a format target that runs multiple format sub-targets.

    Args:
        name: Name of the format target.
        targets: List of format target labels to run.
        **kwargs: Additional arguments passed to sh_binary.
    """
    script_name = name + "_script"

    write_file(
        name = script_name,
        out = name + ".sh",
        content = [
            "#!/bin/bash",
            "set -euo pipefail",
            'cd "$BUILD_WORKSPACE_DIRECTORY"',
            "echo '========================================'",
            "echo 'Running all formatters...'",
            "echo '========================================'",
            "",
        ] + ["bazel run {} && echo".format(t) for t in targets] + [
            "echo '========================================'",
            "echo 'All formatting complete!'",
            "echo '========================================'",
        ],
        is_executable = True,
    )

    native.sh_binary(
        name = name,
        srcs = [":" + script_name],
        **kwargs
    )
