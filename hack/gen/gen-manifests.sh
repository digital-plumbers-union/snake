#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -n "${BUILD_WORKSPACE_DIRECTORY:-}" ]]; then # Running inside bazel
  echo "Updating generated manifests..." >&2
elif ! command -v bazel &>/dev/null; then
  echo "Install bazel via `just tools/bin/bazel`" >&2
  exit 1
else
  echo "This script should be ran via a Bazel target."
  exit 0
fi

go=$(realpath "$1")
controllergen="$(realpath "$2")"
dir="$3"
roleName="$4"
export PATH=$(dirname "$go"):$PATH

# This script should be run via `bazel run //hack:update-crds`
REPO_ROOT=${BUILD_WORKSPACE_DIRECTORY}
cd "${REPO_ROOT}/${dir}"

"$controllergen" \
  crd:trivialVersions=true  \
  rbac:roleName="$roleName" \
  webhook \
  paths="./..." \
  output:crd:artifacts:config=config/crd/bases
