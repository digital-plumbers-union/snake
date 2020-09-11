#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

bin="./bin"
install="bazel run --script_path=$bin"
tools="//hack/tools"

mkdir -p "$bin"

$($install/buildifier $tools:buildifier)
$($install/buildozer $tools:buildozer)
$($install/just $tools:just)
$($install/kustomize $tools:kustomize)
$($install/kubectl $tools:kubectl)
$($install/kubebuilder $tools:kubebuilder)
$($install/kind @io_k8s_sigs_kind//:kind)

