#!/bin/bash
# Format Bazel files with buildifier
set -euo pipefail

# --- begin runfiles.bash initialization ---
if [[ ! -d "${RUNFILES_DIR:-/dev/null}" && ! -f "${RUNFILES_MANIFEST_FILE:-/dev/null}" ]]; then
  if [[ -f "$0.runfiles_manifest" ]]; then
    export RUNFILES_MANIFEST_FILE="$0.runfiles_manifest"
  elif [[ -f "$0.runfiles/MANIFEST" ]]; then
    export RUNFILES_MANIFEST_FILE="$0.runfiles/MANIFEST"
  elif [[ -f "$0.runfiles/bazel_tools/tools/bash/runfiles/runfiles.bash" ]]; then
    export RUNFILES_DIR="$0.runfiles"
  fi
fi
if [[ -f "${RUNFILES_DIR:-/dev/null}/bazel_tools/tools/bash/runfiles/runfiles.bash" ]]; then
  source "${RUNFILES_DIR}/bazel_tools/tools/bash/runfiles/runfiles.bash"
elif [[ -f "${RUNFILES_MANIFEST_FILE:-/dev/null}" ]]; then
  source "$(grep -m1 "^bazel_tools/tools/bash/runfiles/runfiles.bash " "$RUNFILES_MANIFEST_FILE" | cut -d' ' -f2-)"
else
  echo >&2 "ERROR: cannot find @bazel_tools//tools/bash/runfiles:runfiles.bash"
  exit 1
fi
# --- end runfiles.bash initialization ---

# shellcheck source=format-common.sh
source "$(rlocation _main/tools/format-common.sh)"

BUILDIFIER="$(rlocation buildifier_prebuilt~/buildifier/buildifier)"

cd "$BUILD_WORKSPACE_DIRECTORY"

echo "Formatting Bazel files..."
find . -type f \( -name "*.bazel" -o -name "*.bzl" -o -name "BUILD" -o -name "WORKSPACE" \) \
  "${FIND_EXCLUDES[@]}" -print0 \
  | xargs -0 -r "$BUILDIFIER" -mode=fix
echo "Done."
