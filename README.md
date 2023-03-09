# monorepo
Bazel/Go/Protobuf monorepo bootstrap proof of concept

## Setup

These steps needs to be run only once.

### Install asdf

asdf is a package manager that allows pinned tools versions. To install asdf follow the
instructions on https://asdf-vm.com/guide/getting-started.html

### Install Bazel and other tools

After installing asdf and restarting the shell, just run:

```Bash
make install
```

## Build

Although this is not necessary to build with Bazel, it helps to have the GOBIN directory
in your PATH. In order to do that, just run:

```Bash
source setenv
```

or

```Bash
. setenv
```

### Regular build

Runs Gazelle, bazel build and tests

```Bash
make
```

### Run all external tools

Runs Gazelle update dependendencies and a regular build

```Bash
make all
```

### Clean build

Clean bazel, and a full build

```Bash
make clean-build
```

