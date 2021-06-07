#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd -P)"
REPO_ROOT="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"

echo "${SCRIPT_NAME} is running... "

EXIFTOOL_BASE_NAME=Image-ExifTool
EXIFTOOL_VERSION=12.26
EXIFTOOL_ZIP=${EXIFTOOL_BASE_NAME}-${EXIFTOOL_VERSION}.tar.gz

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
  ${SUDO_CMD}rm -rf ${EXIFTOOL_ZIP}
  ${SUDO_CMD}rm -rf ${EXIFTOOL_BASE_NAME}-${EXIFTOOL_VERSION}
}

trap cleanup SIGINT SIGTERM ERR EXIT

SUDO_CMD=""

function downloadExiftool() {
  echo "Gonna to download ${EXIFTOOL_ZIP}"

  wget -O ${EXIFTOOL_ZIP} https://exiftool.org/${EXIFTOOL_ZIP} &&
    ${SUDO_CMD}gzip -dc ${EXIFTOOL_ZIP} | tar -xf - &&
    cd Image-ExifTool-${EXIFTOOL_VERSION} || exit 1 &&
    perl Makefile.PL &&
    ${SUDO_CMD}make install &&
    cd - || exit 1
}

function osxInstall() {
  echo "OSX"

  SUDO_CMD="sudo "

  downloadExiftool
}

function linuxInstall() {
  echo "LINUX"

  DISTRO_ID=$(grep '^ID=' /etc/os-release | sed "s/ID=//")
  echo "${DISTRO_ID}"

  case "${DISTRO_ID}" in
  alpine*)
    SUDO_CMD=""

    downloadExiftool
    ;;
  ubuntu*)
    SUDO_CMD="sudo "
    downloadExiftool
  esac
}

case "$OSTYPE" in
darwin*)
  osxInstall
  ;;
linux*)
  linuxInstall
  ;;
msys* | cygwin*) echo "WINDOWS" ;;
  ## TODO(obalunenko): add windows installation.
*) echo "unknown: $OSTYPE" ;;
esac

echo "Exiftool version: $(exiftool -ver)"

echo "${SCRIPT_NAME} done."
