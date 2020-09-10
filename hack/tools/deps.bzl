load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

def install():
    install_kubebuilder()
    install_bazel_tools()
    install_kustomize()

def install_bazel_tools():
    """Install additional tools related to Bazel

    These are not installed via tools.go because of their dependency on protobuf
    """
    http_file(
        name = "buildozer_osx",
        sha256 = "81cb08a5d73a41643e07b163adf1a2fcc4f30d9c9a0f5f1f2b258b5ba94c9bbb",
        executable = 1,
        urls = ["https://github.com/bazelbuild/buildtools/releases/download/3.4.0/buildozer.mac"],
    )

    http_file(
        name = "buildozer_linux",
        executable = 1,
        urls = ["https://github.com/bazelbuild/buildtools/releases/download/3.4.0/buildozer"],
    )

    http_file(
        name = "buildifier_osx",
        sha256 = "3c30fcddfea8b515fff75127788c16dca5d901873ec4cf2102225cccbffc1702",
        executable = 1,
        urls = ["https://github.com/bazelbuild/buildtools/releases/download/3.4.0/buildifier.mac"],
    )

    http_file(
        name = "buildifier_linux",
        sha256 = "5d47f5f452bace65686448180ff63b4a6aaa0fb0ce0fe69976888fa4d8606940",
        executable = 1,
        urls = ["https://github.com/bazelbuild/buildtools/releases/download/3.4.0/buildifier"],
    )

def install_kubebuilder():
    http_archive(
        name = "kubebuilder_linux",
        build_file_content = """
package(default_visibility = ["//visibility:public"])
load("@rules_pkg//:pkg.bzl", "pkg_tar")
pkg_tar(
  name = "tar",
  # contents of this zip are kubebuilder_2.2.0_linux_amd64/bin/*, we want all
  srcs = glob(["kubebuilder_2.2.0_linux_amd64/bin/*"]),
  extension = "tar.gz",
  mode = "755",
  package_dir = "/usr/local/kubebuilder/",
  strip_prefix = "kubebuilder_2.2.0_linux_amd64/"
)
filegroup(
  name = "file",
  srcs = glob(["kubebuilder_2.2.0_linux_amd64/bin/*"]),
  visibility = ["//visibility:public"]
)
""",
        sha256 = "9ef35a4a4e92408f7606f1dd1e68fe986fa222a88d34e40ecc07b6ffffcc8c12",
        urls = ["https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.2.0/kubebuilder_2.2.0_linux_amd64.tar.gz"],
    )
    http_archive(
        name = "kubebuilder_osx",
        build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
  name = "file",
  srcs = glob(["kubebuilder_2.2.0_darwin_amd64/bin/*"], ["kubebuilder_2.2.0_darwin_amd64/bin/kubectl"]),
  visibility = ["//visibility:public"]
)
""",
        sha256 = "5ccb9803d391e819b606b0c702610093619ad08e429ae34401b3e4d448dd2553",
        urls = ["https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.2.0/kubebuilder_2.2.0_darwin_amd64.tar.gz"],
    )

def install_kustomize():
    http_archive(
        name = "kustomize_linux",
        build_file_content = """
package(default_visibility = ["//visibility:public"])
load("@rules_pkg//:pkg.bzl", "pkg_tar")
pkg_tar(
  name = "tar",
  srcs = ["kustomize"],
  extension = "tar.gz",
  mode = "755",
  package_dir = "/usr/local/bin",
  # strip prefix so dir isnt flattened, within `http_archive`, that value is "."
  # https://github.com/bazelbuild/rules_docker/issues/317
  strip_prefix = "."
)
filegroup(
  name = "file",
  srcs = ["kustomize"],
  visibility = ["//visibility:public"]
)
""",
        sha256 = "5cdeb2af81090ad428e3a94b39779b3e477e2bc946be1fe28714d1ca28502f6a",
        urls = ["https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv3.5.4/kustomize_v3.5.4_linux_amd64.tar.gz"],
    )
    http_archive(
        name = "kustomize_osx",
        build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
  name = "file",
  srcs = ["kustomize"],
  visibility = ["//visibility:public"]
)
""",
        sha256 = "9215c140593537b30e83f14277dee8a2adea9bb44825a8ed98a6c12a82fb2ea6",
        urls = ["https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv3.5.4/kustomize_v3.5.4_darwin_amd64.tar.gz"],
    )

