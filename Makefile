# Makefile helper around bazel commands
#
# Makefile targets:
# - make or make default: run default code generation tools, gazelle, bazel build and bazel build
# - make all: run all code generation tools (including gazelle-update-modules) and make default
# - make clean-build: cleans the build and the run make all

export GO111MODULE := on

.DEFAULT_GOAL := default

.PHONY: default
default: build test

.PHONY: build
build: gazelle protolink test
	bazel build //...

.PHONY: all
all: gazelle-update-repos build

.PHONY: clean-build
clean-build: clean all

.PHONY: install
install:
	asdf update
	-cat .tool-versions | awk '{print $$1}' | xargs --no-run-if-empty -L 1 asdf plugin add
	asdf install
	@echo Download go.mod dependencies
	@go mod download
	go install github.com/bazelbuild/buildtools/buildifier@latest

.PHONY: test
test: build
	bazel test --test_verbose_timeout_warnings //...

.PHONY: gazelle
gazelle:
	bazel run //:gazelle -- update

.PHONY: gazelle-update-repos
gazelle-update-repos:
	bazel run //:gazelle-update-repos

.PHONY: protolink
protolink:
	bazel query 'kind(".*_proto_link", //...)' | xargs -L 1 bazel run

.PHONY: clean
clean:
	bazel clean --expunge
