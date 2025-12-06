"Define linter aspects for the monorepo"

load("@aspect_rules_lint//lint:lint_test.bzl", "lint_test")
load("@aspect_rules_lint//lint:ruff.bzl", "lint_ruff_aspect")
load("@aspect_rules_lint//lint:shellcheck.bzl", "lint_shellcheck_aspect")
load("//tools/lint:golangci_lint.bzl", "lint_golangci_lint_aspect")

# Go linting with golangci-lint v2
golangci_lint = lint_golangci_lint_aspect(
    binary = "@multitool//tools/golangci-lint",
    config = Label("//:.golangci.yml"),
)

golangci_lint_test = lint_test(aspect = golangci_lint)

# Python linting with ruff
ruff = lint_ruff_aspect(
    binary = "@multitool//tools/ruff",
    configs = [
        Label("//:.ruff.toml"),
    ],
)

ruff_test = lint_test(aspect = ruff)

# Shell script linting with shellcheck
shellcheck = lint_shellcheck_aspect(
    binary = "@multitool//tools/shellcheck",
    config = Label("//:.shellcheckrc"),
)

shellcheck_test = lint_test(aspect = shellcheck)
