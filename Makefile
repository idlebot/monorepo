export GO111MODULE := on

.DEFAULT_GOAL := all

.PHONY: all
all: gazelle gazelle-update-repos build test

.PHONY: build
build:
	bazel run //hellogrpc/greeter/v1:v1_go_proto_link
	bazel build //...

.PHONY: test
test: build
	bazel test //...

.PHONY: gazelle
gazelle:
	bazel run //:gazelle -- update

.PHONY: gazelle-update-repos
gazelle-update-repos:
	bazel run //:gazelle-update-repos

.PHONY: clean
clean:
	bazel clean --expunge
