# monorepo
Multi language Bazel monorepo proof of concept

# Development Environment Setup

This repository provides a standardized development environment setup using Bazel.

## Getting Started

### Prerequisites

- Git
- curl

### Environment Setup

The `setenv` script creates an isolated shell environment with the correct PATH and tools configured for this repository. This ensures all developers use consistent tooling regardless of their local setup.

To use it:

```bash
./setenv
```

This will:
1. Add the repository's `bin/` directory to your PATH
2. Download and configure bazelisk (a version manager for Bazel)
3. Launch a new shell with these configurations

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
│   └── go/        # Go source code
│       └── helloworld/  # Example Go program
└── BUILD.bazel    # Root Bazel build file
```
