#!/bin/bash
set -euo pipefail

# ==============================================================================
# DEVSHELL - Development Shell with Bazel-managed Tools
# ==============================================================================
#
# Usage: ./devshell.sh
#
# This script creates a development shell with Bazel-managed tools available
# directly in PATH (go, python3, protoc, buildifier, etc.)
#
# ┌─────────────────────────────────────────────────────────────────────────────┐
# │ HOW TO ADD A NEW TOOL                                                       │
# ├─────────────────────────────────────────────────────────────────────────────┤
# │ 1. Add the Bazel target to TOOLS array below                                │
# │ 2. Add wrapper generator function if special handling needed                │
# │ 3. Run: ./devshell.sh                                                       │
# └─────────────────────────────────────────────────────────────────────────────┘
#
# ==============================================================================

# ------------------------------------------------------------------------------
# CONFIGURATION
# ------------------------------------------------------------------------------

# Tools to build and expose in PATH
# Format: "bazel_target:binary_name"
TOOLS=(
  "@go_default_sdk//:bin/go:go"
  "@com_google_protobuf//:protoc:protoc"
  "@python_3_14//:python3:python3"
  "@buildifier_prebuilt//:buildifier:buildifier"
  "//tools:shfmt:shfmt"
)

# ------------------------------------------------------------------------------
# PATH SETUP
# ------------------------------------------------------------------------------

# Get repo root (same as script location)
REPO_ROOT="$(cd "$(dirname "$0")" && pwd)"
BIN_DIR="${REPO_ROOT}/bin"

# ------------------------------------------------------------------------------
# ENVIRONMENT CHECK
# ------------------------------------------------------------------------------

if [[ -n "${DEVSHELL_ACTIVE:-}" ]]; then
  echo "Error: Already in devshell. Type 'exit' to leave first."
  exit 1
fi

# ------------------------------------------------------------------------------
# TOOL WRAPPER GENERATORS
# ------------------------------------------------------------------------------

# Default wrapper - just executes the binary
generate_default_wrapper() {
  local binary_path="$1"
  local runfiles_dir="$2"

  cat << EOF
#!/bin/bash
export RUNFILES_DIR="$runfiles_dir"
exec "$binary_path" "\$@"
EOF
}

# Go wrapper - needs GOROOT and GOTOOLCHAIN handling
generate_go_wrapper() {
  local binary_path="$1"
  local runfiles_dir="$2"

  # GOROOT is the SDK root (parent of bin/)
  local goroot
  goroot="$(dirname "$(dirname "$binary_path")")"

  cat << EOF
#!/bin/bash
export GOROOT="$goroot"
export GOTOOLCHAIN=local
exec "$binary_path" "\$@"
EOF
}

# Python wrapper - ensure clean environment
generate_python3_wrapper() {
  local binary_path="$1"
  local runfiles_dir="$2"

  cat << EOF
#!/bin/bash
export RUNFILES_DIR="$runfiles_dir"
exec "$binary_path" "\$@"
EOF
}

# ------------------------------------------------------------------------------
# BUILD AND INSTALL TOOLS
# ------------------------------------------------------------------------------

build_tools() {
  local targets=()
  for tool in "${TOOLS[@]}"; do
    local target="${tool%:*}"
    targets+=("$target")
  done

  echo ""
  echo "🔨 Building tools..."
  bazel build "${targets[@]}"
}

install_tool_wrappers() {
  echo ""
  echo "📝 Creating tool wrappers..."

  local output_base
  output_base=$(bazel info output_base)

  for tool in "${TOOLS[@]}"; do
    local target="${tool%:*}"
    local binary_name="${tool##*:}"

    # Get binary path from bazel
    local binary_path
    binary_path=$(bazel cquery --output=files "$target" 2> /dev/null | grep "${binary_name}$" | head -1 || true)

    if [[ -z "$binary_path" ]]; then
      echo "  ✗ $binary_name: binary not found"
      continue
    fi

    # Resolve full path
    local full_path runfiles_dir
    if [[ "$binary_path" == external/* ]]; then
      full_path="$output_base/$binary_path"
      runfiles_dir="$output_base/$(dirname "$binary_path").runfiles"
    else
      full_path="$(realpath "$binary_path")"
      runfiles_dir="$(dirname "$full_path").runfiles"
    fi

    # Generate appropriate wrapper
    local wrapper_func="generate_default_wrapper"
    if declare -f "generate_${binary_name}_wrapper" > /dev/null; then
      wrapper_func="generate_${binary_name}_wrapper"
    fi

    "$wrapper_func" "$full_path" "$runfiles_dir" > "$BIN_DIR/$binary_name"
    chmod +x "$BIN_DIR/$binary_name"

    # Show version if possible
    local version
    version=$("$BIN_DIR/$binary_name" --version 2> /dev/null | head -1 || echo "installed")
    echo "  ✓ $binary_name: $version"
  done
}

# ------------------------------------------------------------------------------
# PYTHON VENV SETUP
# ------------------------------------------------------------------------------

setup_python_venv() {
  echo ""
  echo "🐍 Setting up Python environment..."

  local venv_dir="$REPO_ROOT/.venv"

  if [[ ! -f "$venv_dir/bin/activate" ]]; then
    echo "  Creating virtual environment..."
    "$BIN_DIR/python3" -m venv "$venv_dir"
  fi

  # shellcheck disable=SC1091
  source "$venv_dir/bin/activate"

  if [[ -f "$REPO_ROOT/requirements.txt" ]]; then
    pip install -q -r "$REPO_ROOT/requirements.txt"
  fi

  echo "  ✓ Virtual environment ready"
}

# ------------------------------------------------------------------------------
# LAUNCH SHELL
# ------------------------------------------------------------------------------

launch_shell() {
  local new_path="$REPO_ROOT/.venv/bin:$BIN_DIR:$PATH"

  echo ""
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo "  DEVSHELL ready! Tools available: ${TOOLS[*]##*:}"
  echo "  Type 'exit' to leave"
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo ""

  # Create a temporary rc file for the shell
  local rc_file
  rc_file=$(mktemp)

  # Create shell-specific rc file with custom prompt
  case "${SHELL:-/bin/bash}" in
    */zsh)
      local zsh_dir
      zsh_dir=$(mktemp -d)
      cat > "$zsh_dir/.zshrc" << 'ZSHRC'
[[ -f ~/.zshrc ]] && source ~/.zshrc
PROMPT='%F{cyan}[devshell]%f %F{blue}%1~%f %# '
ZSHRC
      exec env \
        PATH="$new_path" \
        DEVSHELL_ACTIVE=1 \
        VIRTUAL_ENV="$REPO_ROOT/.venv" \
        PYTHONPATH="$REPO_ROOT/src/python" \
        GOPATH="$REPO_ROOT/.go" \
        ZDOTDIR="$zsh_dir" \
        zsh
      ;;
    *)
      cat > "$rc_file" << 'BASHRC'
[[ -f ~/.bashrc ]] && source ~/.bashrc
PS1='\[\033[1;36m\][devshell]\[\033[0m\] \[\033[1;34m\]\W\[\033[0m\] \$ '
BASHRC
      exec env \
        PATH="$new_path" \
        DEVSHELL_ACTIVE=1 \
        VIRTUAL_ENV="$REPO_ROOT/.venv" \
        PYTHONPATH="$REPO_ROOT/src/python" \
        GOPATH="$REPO_ROOT/.go" \
        bash --rcfile "$rc_file"
      ;;
  esac
}

# ------------------------------------------------------------------------------
# MAIN
# ------------------------------------------------------------------------------

main() {
  echo ""
  echo "┌─────────────────────────────────────────────────────────────────────┐"
  echo "│  DEVSHELL - Bazel Development Environment                           │"
  echo "└─────────────────────────────────────────────────────────────────────┘"
  echo ""

  cd "$REPO_ROOT"
  mkdir -p "$BIN_DIR"

  build_tools
  install_tool_wrappers
  setup_python_venv
  launch_shell
}

main "$@"
