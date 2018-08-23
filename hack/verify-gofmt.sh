#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

diff=$(find . -name "*.go" | grep -v "\/vendor\/" | xargs gofmt -s -d 2>&1)
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  echo
  echo "Please run hack/update-gofmt.sh"
  exit 1
fi

echo "done"