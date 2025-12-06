"""API for declaring a golangci-lint aspect that visits go_library, go_binary, go_test rules.

Usage in tools/lint/linters.bzl:

```starlark
load("//tools/lint:golangci_lint.bzl", "lint_golangci_lint_aspect")

golangci_lint = lint_golangci_lint_aspect(
    binary = "@multitool//tools/golangci-lint",
    config = Label("//:.golangci.yml"),
)
```
"""

load("@aspect_rules_lint//lint/private:lint_aspect.bzl", "LintOptionsInfo", "filter_srcs", "noop_lint_action", "output_files", "should_visit")

_MNEMONIC = "AspectRulesLintGolangciLint"

def _golangci_lint_action(ctx, executable, srcs, config, go_sdk, stdout, exit_code = None):
    """Run golangci-lint as an action under Bazel.

    Args:
        ctx: Bazel Rule or Aspect evaluation context
        executable: label of the golangci-lint program
        srcs: Go files to be linted
        config: label of the .golangci.yml file
        go_sdk: the Go SDK to use
        stdout: output file containing stdout of golangci-lint
        exit_code: output file containing golangci-lint exit code.
            If None, then fail the build when golangci-lint exits non-zero.
    """
    inputs = srcs + [config]
    tools = [executable]

    # Get Go SDK paths
    go_root = go_sdk.root_file.dirname

    # Add Go SDK to inputs (libs, srcs, headers, tools are depsets, so use to_list())
    inputs = inputs + go_sdk.libs.to_list() + go_sdk.srcs.to_list() + go_sdk.headers.to_list() + go_sdk.tools.to_list() + [go_sdk.go, go_sdk.root_file]

    args = ctx.actions.args()
    args.add("run")
    args.add("--config", config)
    args.add("--output.text.colors")
    args.add("--issues-exit-code", "1")
    args.add_all(srcs)

    outputs = [stdout]

    # Set up environment to find Go and cache
    # Use $PWD to make relative paths absolute in the sandbox
    env_setup = "export PATH=$PWD/{go_bin_dir}:$PATH && export GOROOT=$PWD/{go_root} && export HOME=$PWD && export GOLANGCI_LINT_CACHE=$PWD/.golangci-cache && mkdir -p $GOLANGCI_LINT_CACHE && ".format(
        go_bin_dir = go_sdk.go.dirname,
        go_root = go_root,
    )

    if exit_code:
        command = env_setup + "{golangci_lint} $@ >{stdout} 2>&1; echo $? >" + exit_code.path
        outputs.append(exit_code)
    else:
        command = env_setup + "{golangci_lint} $@ >{stdout} 2>&1 && touch {stdout}"

    ctx.actions.run_shell(
        inputs = inputs,
        outputs = outputs,
        command = command.format(
            golangci_lint = executable.path,
            stdout = stdout.path,
        ),
        arguments = [args],
        mnemonic = _MNEMONIC,
        progress_message = "Linting %{label} with golangci-lint",
        tools = tools,
    )

def _golangci_lint_aspect_impl(target, ctx):
    if not should_visit(ctx.rule, ctx.attr._rule_kinds):
        return []

    files_to_lint = filter_srcs(ctx.rule)
    outputs, info = output_files(_MNEMONIC, target, ctx)

    if len(files_to_lint) == 0:
        noop_lint_action(ctx, outputs)
        return [info]

    # Get Go SDK from toolchain
    go_toolchain = ctx.toolchains["@rules_go//go:toolchain"]
    go_sdk = go_toolchain.sdk

    _golangci_lint_action(
        ctx,
        ctx.executable._golangci_lint,
        files_to_lint,
        ctx.file._config_file,
        go_sdk,
        outputs.human.out,
        outputs.human.exit_code,
    )

    # golangci-lint doesn't have separate machine-readable output easily,
    # so we just copy the human output for now
    ctx.actions.run_shell(
        inputs = [outputs.human.out],
        outputs = [outputs.machine.out],
        command = "cp {input} {output}".format(
            input = outputs.human.out.path,
            output = outputs.machine.out.path,
        ),
        mnemonic = _MNEMONIC + "CopyReport",
    )

    # Copy exit code too
    ctx.actions.run_shell(
        inputs = [outputs.human.exit_code],
        outputs = [outputs.machine.exit_code],
        command = "cp {input} {output}".format(
            input = outputs.human.exit_code.path,
            output = outputs.machine.exit_code.path,
        ),
        mnemonic = _MNEMONIC + "CopyExitCode",
    )

    return [info]

def lint_golangci_lint_aspect(binary, config, rule_kinds = ["go_library", "go_binary", "go_test"]):
    """A factory function to create a linter aspect for golangci-lint.

    Args:
        binary: a golangci-lint executable.
        config: the .golangci.yml config file
        rule_kinds: which rule kinds to lint (default: go_library, go_binary, go_test)

    Returns:
        An aspect that runs golangci-lint on Go targets.
    """
    return aspect(
        implementation = _golangci_lint_aspect_impl,
        attrs = {
            "_options": attr.label(
                default = "@aspect_rules_lint//lint:options",
                providers = [LintOptionsInfo],
            ),
            "_golangci_lint": attr.label(
                default = binary,
                executable = True,
                cfg = "exec",
            ),
            "_config_file": attr.label(
                default = config,
                allow_single_file = True,
            ),
            "_rule_kinds": attr.string_list(
                default = rule_kinds,
            ),
        },
        toolchains = ["@rules_go//go:toolchain"],
    )
