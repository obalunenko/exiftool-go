#!/bin/bash

set -Eeuo pipefail

function cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  echo "cleanup running"
  #rm -rf coverage*.out.tmp
}

trap cleanup SIGINT SIGTERM ERR EXIT

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd -P)"
REPO_ROOT="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"

echo "${SCRIPT_NAME} is running... "

rm -rf coverage
mkdir -p coverage

go test --count=1 -tags=integration_test -coverprofile ./coverage/integration.cov -covermode=atomic ./...
go test --count=1 -coverprofile ./coverage/unit.cov -covermode=atomic ./...


{
echo "mode: atomic"
tail -q -n +2 ./coverage/*.cov
} >> ./coverage/full.cov

gocov convert "${COVER_DIR}/full.cov" >"${COVER_DIR}/full.json"

echo "${SCRIPT_NAME} done."
