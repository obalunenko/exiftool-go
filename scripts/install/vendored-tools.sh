#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

cd "${TOOLS_DIR}" || exit 1

SUDO_CMD=""

function osxInstall() {
  echo "OSX"

  SUDO_CMD="sudo "

  install_deps
}

function linuxInstall() {
  echo "LINUX"

  DISTRO_ID=$(grep '^ID=' /etc/os-release | sed "s/ID=//")
  echo "${DISTRO_ID}"

  case "${DISTRO_ID}" in
  alpine*)
    SUDO_CMD=""

    install_deps
    ;;
  ubuntu*)
    SUDO_CMD="sudo "
    install_deps
    ;;
  esac

  cleanup
}

function check_status() {
  # first param is error message to print in case of error
  if [ $? -ne 0 ]; then
    if [ -n "$1" ]; then
      echo "$1"
    fi

    # Exit 255 to pass signal to xargs to abort process with code 1, in other cases xargs will complete with 0.
    exit 255
  fi
}

function install_dep() {
  dep=$1

  bin_out=$GOBIN/$(echo $dep | awk 'BEGIN { FS="/" } {print $NF}')

  echo "[INFO]: Going to build ${dep} - ${bin_out}"

  ${SUDO_CMD}go build -mod=vendor -o "${bin_out}" "${dep}"

  check_status "[FAIL]: build [${dep}] failed!"

  echo "[SUCCESS]: build [${dep}] finished."
}

export -f install_dep
export -f check_status

function install_deps() {
  tools_module="$(go list -m)"

  go list -e -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
   xargs -n 1 -P 0 -I {} bash -c 'install_dep "$@"' _ {}
}

OS_TYPE="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS_TYPE" in
darwin*)
  osxInstall
  ;;
linux*)
  linuxInstall
  ;;
msys* | cygwin*) echo "WINDOWS" ;;
  ## TODO(o.balunenko): add windows installation.
*) echo "unknown: $OS_TYPE" ;;
esac