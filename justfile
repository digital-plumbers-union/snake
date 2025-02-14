# installs required tooling for this repository, ran by default if no recipe is provided
setup:
  hack/tools/install.sh

# cleans workspace, deletes kind clusters
clean: kind-down
  bazel clean --expunge
  rm -rf bin/

# cleans workspace, reinstalls tools
refresh: clean setup

# update BUILD files & build
build: gazelle
  bazel build //...

# update BUILD files & test
test: gazelle
  bazel test //...

# update BUILD files
gazelle:
  bazel run //:gazelle

# update external go deps in bazel
update-go-deps:
  go mod tidy
  bazel run //:gazelle -- update-repos -from_file=go.mod -prune=true --build_file_generation=on --build_file_proto_mode=disable_global

# run basic formatting + linting check against code
check:
  just bazel-style check

# run formatting/style updates that can be automated
fix:
  just bazel-style

# uses buildifier to format.  pass mode=check to check without fixing
bazel-style mode="fix":
  if test -e bin/buildifier; then bin/buildifier --mode {{mode}} -r `pwd`; else bazel run //hack/tools:buildifier -- --mode {{mode}} -r `pwd`; fi

# stand up kind development cluster and install pre-reqs
kind-up:
  #!/usr/bin/env sh
  if kind get clusters | grep "snake-dev"; then
    echo "snake-dev cluster already created"
  else
    kind create cluster --name snake-dev --kubeconfig `pwd`/kind-kubeconfig.yaml
    # install pre-reqs that we dont want to install every time we refresh our manifests
    kubectl apply -f hack/cert-manager.yaml
  fi

# tear down kind development cluster
kind-down:
  if kind get clusters | grep "snake-dev"; then kind delete cluster --name snake-dev; else echo "No snake-dev cluster to delete"; fi

# run all manifest generation targets
manifests:
  #!/usr/bin/env sh
  GENTARGETS=$(bazel query "kind(sh_binary, attr('generator_name', 'gen_manifests', //...))")
  for target in "${GENTARGETS[@]}"
  do
    bazel run "$target"
  done

# installs manifests for all components to active kubectl context
install: manifests
  kustomize build scheduler/config | kubectl apply -f -

# uninstalls manifests
uninstall:
  kustomize build scheduler/config | kubectl delete -f -
