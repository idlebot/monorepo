#!/bin/bash
# Format shell scripts with shfmt
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

SHFMT="$(rlocation _main~prebuilt_tools~shfmt_prebuilt/shfmt)"

cd "$BUILD_WORKSPACE_DIRECTORY"

# Google Shell Style Guide formatting:
#   -i 2  : 2 space indentation
#   -bn   : binary ops (&&, |) may start a line
#   -ci   : switch case bodies are indented
#   -sr   : redirect operators followed by a space
SHFMT_FLAGS="-i 2 -bn -ci -sr"

echo "Formatting shell scripts (Google style)..."
find . -type f -name "*.sh" "${FIND_EXCLUDES[@]}" -print0 \
  | xargs -0 -r "$SHFMT" $SHFMT_FLAGS -w
echo "Done."
