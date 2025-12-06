#!/bin/bash
# Run all linters via Bazel aspects
# Usage: bazel run //:lint [target_pattern]

set -euo pipefail

cd "$BUILD_WORKSPACE_DIRECTORY"

TARGET_PATTERN="${1:-//...}"

echo "========================================"
echo "Running linters on $TARGET_PATTERN"
echo "========================================"

# Define the aspects to run
ASPECTS=(
  "//tools/lint:linters.bzl%ruff"
  "//tools/lint:linters.bzl%shellcheck"
)

# Join aspects with comma
ASPECT_ARG=$(
  IFS=,
  echo "${ASPECTS[*]}"
)

# Run bazel build with lint aspects
# Output groups: rules_lint_report for human-readable, rules_lint_patch for fixes
bazel build "$TARGET_PATTERN" \
  --aspects="$ASPECT_ARG" \
  --output_groups=rules_lint_report \
  --keep_going \
  2>&1

# Print any lint reports with actual issues (not just "All checks passed!")
echo ""
echo "========================================"
echo "Lint Results:"
echo "========================================"

found_issues=0
while IFS= read -r report; do
  if [[ -s "$report" ]]; then
    # Skip reports that only contain "All checks passed!" or are empty
    content=$(cat "$report")
    if [[ -n "$content" && "$content" != "All checks passed!" ]]; then
      found_issues=1
      echo ""
      echo "$content"
    fi
  fi
done < <(find -L bazel-bin -name "*AspectRulesLint*.report" -type f 2> /dev/null)

echo "========================================"
if [[ $found_issues -eq 0 ]]; then
  echo "Linting complete! No issues found."
  exit 0
else
  echo "Linting complete! Issues found."
  exit 1
fi
