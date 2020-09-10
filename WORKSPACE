workspace(
    name = "snake",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

################################################################################
# BAZEL RULE / TOOLCHAIN SETUP
################################################################################

# download `io_bazel_rules_go` up front to ensure all of our other rulesets
# leverage the same version, see related issue:
# https://github.com/bazelbuild/rules_go/issues/2398#issuecomment-597139571
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "08c3cd71857d58af3cda759112437d9e63339ac9c6e0042add43f4d94caf632d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
    ],
)

#########################################
# DOCKER
#########################################

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "dc97fccceacd4c6be14e800b2a00693d5e8d07f69ee187babfd04a80a9f8e250",
    strip_prefix = "rules_docker-0.14.1",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.1/rules_docker-v0.14.1.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

# configures the docker toolchain, https://github.com/nlopezgi/rules_docker/blob/master/toolchains/docker/readme.md#how-to-use-the-docker-toolchain
container_repositories()

#########################################
# GOLANG
#########################################

# set up `io_bazel_rules_go` imported at top of this file
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# gazell generates BUILD files for go
http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# only set up dependencies once we have imported everything that could
# possibly be overridden: see "overriding dependencies" here:
# https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id9
go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

################################################################################
# EXTERNAL DEPENDENCIES
################################################################################

##########################################################
# TOOLS
##########################################################

load("//hack/tools:deps.bzl", install_tools = "install")

install_tools()

##########################################################
# GO DEPENDENCIES
##########################################################

load(":deps.bzl", "go")

# this function is generated using gazelle, we load and execute it here to
# reduce WORKSPACE file size
go()

# pull in go_image deps
# see: https://github.com/bazelbuild/rules_docker#go_image

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
