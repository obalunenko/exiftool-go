#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"
COVER_DIR=${REPO_ROOT}/coverage

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

export GO111MODULE=on

rm -rf "${COVER_DIR}"
mkdir -p "${COVER_DIR}"

go test --count=1 -tags=integration_test -coverprofile "${COVER_DIR}/unit.cov" -covermode=atomic -json ./... >tests-report.json

echo "${SCRIPT_NAME} is done... "
