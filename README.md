# Monorepo

Multi-language Bazel monorepo template supporting Go, Python, and C++ projects.

## Development Environment Setup

This repository provides a standardized development environment using Bazel, ensuring consistent tooling across all developers.

### Prerequisites

- Git
- curl

### Environment Setup

The repository includes an environment setup script that creates an isolated shell with all necessary tools and configurations:

```sh
source ./env
```

or

```sh
. ./env
```


This will:
1. Add the repository's `bin/` directory to your PATH
2. Download and configure bazelisk (a version manager for Bazel)
3. Provide hermetic Bazel versions for each supported language tools

To exit the environment, simply type `exit` or press `Ctrl+D`.

### Building and Running

This repository uses [Bazel](https://bazel.build/) as its build system. 

#### Updating BUILD Files

We use [Gazelle](https://github.com/bazelbuild/bazel-gazelle) to automatically manage Bazel BUILD files for Go code. To update BUILD files after adding or modifying Go code:

```bash
bazel run //:gazelle
```

#### Building Projects

To build a specific target:
```bash
bazel build //path/to/target
```

To run a target:
```bash
bazel run //path/to/target
```

For example, to run the hello world program:
```bash
bazel run //src/go/helloworld
```

## Project Structure

```
.
├── bin/           # Development tools and scripts
├── src/          
│   ├── go/        # Go source code
│   │   └── helloworld/  # Example Go program
│   └── python/    # Python source code
│       └── helloworld/  # Example Python program
└── BUILD.bazel    # Root Bazel build file
```
