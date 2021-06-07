#!/usr/bin/env bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd -P)"
REPO_ROOT="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"

echo "${SCRIPT_NAME} is running... "

GO_FILES=$(find . -type f -name '*.go' | grep -v 'vendor' | grep -v '.git')

if [[ -f "$(go env GOPATH)/bin/golines" ]] || [[ -f "/usr/local/bin/golines" ]]; then
golines -w -m 120 --no-reformat-tags ${GO_FILES}
else
  printf "Cannot check golines, please run:
    make install-tools \n"
  exit 1
fi

echo "${SCRIPT_NAME} done."
