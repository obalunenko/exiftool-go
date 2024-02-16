#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

EXIFTOOL_BASE_NAME=exiftool
EXIFTOOL_VERSION=${1} # 12.52
EXIFTOOL_ZIP=${EXIFTOOL_BASE_NAME}-${EXIFTOOL_VERSION}.tar.gz

SUDO_CMD=""

cleanup() {
  echo "cleanup running"
  ${SUDO_CMD}rm -rf ${EXIFTOOL_ZIP}
  ${SUDO_CMD}rm -rf ${EXIFTOOL_BASE_NAME}-${EXIFTOOL_VERSION}
}

# https://github.com/exiftool/exiftool/archive/refs/tags/12.28.tar.gz

downloadExiftool() {
  echo "Gonna to download ${EXIFTOOL_ZIP}"

  wget --no-check-certificate -O ${EXIFTOOL_ZIP} https://github.com/exiftool/exiftool/archive/refs/tags/${EXIFTOOL_VERSION}.tar.gz &&
    tar -xvf ${EXIFTOOL_ZIP} &&
    cd ${EXIFTOOL_BASE_NAME}-${EXIFTOOL_VERSION} || exit 1 &&
    perl Makefile.PL &&
    ${SUDO_CMD}make install &&
    cd - || exit 1
}

osxInstall() {
  echo "OSX"

  SUDO_CMD="sudo "

  downloadExiftool

  cleanup
}

linuxInstall() {
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
    ;;
  esac

  cleanup
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

echo "Exiftool version: $(exiftool -ver)"

echo "${SCRIPT_NAME} done."
