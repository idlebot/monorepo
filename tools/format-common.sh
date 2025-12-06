#!/bin/bash
# Common configuration for format scripts
# Source this file after runfiles initialization

# Common directories to exclude from formatting
# Usage: find . -type f -name "*.ext" "${FIND_EXCLUDES[@]}" -print0
FIND_EXCLUDES=(
  -not -path "./.git/*"
  -not -path "./bazel-*"
  -not -path "./.venv/*"
  -not -path "./.go/*"
)
