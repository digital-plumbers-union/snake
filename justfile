################################################################################
# variables
################################################################################

# Project Variables
# Use `env_var` as much as possible to leverage `.env` as single source of truth

# commit := `git rev-parse HEAD`
# branch := `git rev-parse --abbrev-ref HEAD`

################################################################################
# commands
################################################################################

################################################################################
# dependency management recipes
################################################################################

setup:
  hack/tools/install.sh

# NOTE: first command is default command
# (i.e., what happens when you run `just` with no recipe)
# update BUILD files & build
build: gazelle
  bazel build //...

# update BUILD files & test
test: gazelle
  bazel test //...

# update BUILD files
gazelle:
  bazel run //:gazelle

# runs as container
run-container:
  bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //scheduler:container

# publish container to docker hub
publish-container:
  bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd:push-container

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
