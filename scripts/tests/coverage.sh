#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"
COVER_DIR=${REPO_ROOT}/coverage

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

checkInstalled gocov

rm -rf "${COVER_DIR}"
mkdir -p "${COVER_DIR}"

go test --count=1 -tags=integration_test -coverprofile ./coverage/integration.cov -covermode=atomic ./...
go test --count=1 -coverprofile ./coverage/unit.cov -covermode=atomic ./...

{
echo "mode: atomic"
tail -q -n +2 ./coverage/*.cov
} >> ./coverage/full.cov

gocov convert "${COVER_DIR}/full.cov" >"${COVER_DIR}/full.json"

gocov report "${COVER_DIR}/full.json"

echo "${SCRIPT_NAME} done."
