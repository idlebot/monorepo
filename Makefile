# Makefile helper around bazel commands
#
# Makefile targets:
# - make or make default: run default code generation tools, gazelle, bazel build and bazel build
# - make all: run all code generation tools (including gazelle-update-modules) and make default
# - make clean-build: cleans the build and the run make all

export GO111MODULE := on

.DEFAULT_GOAL := default

.PHONY: default
default: build

.PHONY: build
build: gazelle
	bazel run --jobs=6 //:buildifier
	bazel build --jobs=6 //...

.PHONY: all
all: gazelle-update-repos build

.PHONY: clean-build
clean-build: clean all

.PHONY: install
install:
	-asdf update
	-cat .tool-versions | awk '{print $$1}' | xargs -L 1 asdf plugin add
	asdf install
	@echo Download go.mod dependencies
	@go mod download
	go install github.com/bazelbuild/buildtools/buildifier@latest

.PHONY: test
test: build
	bazel test --test_verbose_timeout_warnings //...

.PHONY: gazelle
gazelle:
	bazel run --jobs=6 //:gazelle -- update

.PHONY: gazelle-update-repos
gazelle-update-repos:
	bazel run --jobs=6 //:gazelle-update-repos

.PHONY: clean
clean:
	bazel clean --expunge
