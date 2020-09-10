package(default_visibility = ["//visibility:public"])

genrule(
    name = "fetch_kubectl",
    srcs = select({
        ":darwin": ["@kubectl_osx//file"],
        ":k8": ["@kubectl_linux//file"],
    }),
    outs = ["kubectl"],
    cmd = "cp $(SRCS) $@",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

genrule(
    name = "fetch_buildifier",
    srcs = select({
        ":darwin": ["@buildifier_osx//file"],
        ":k8": ["@buildifier_linux//file"],
    }),
    outs = ["buildifier"],
    cmd = "cp $(SRCS) $@",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

genrule(
    name = "fetch_buildozer",
    srcs = select({
        ":darwin": ["@buildozer_osx//file"],
        ":k8": ["@buildozer_linux//file"],
    }),
    outs = ["buildozer"],
    cmd = "cp $(SRCS) $@",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

genrule(
    name = "fetch_kubebuilder",
    srcs = select({
        ":darwin": ["@kubebuilder_osx//:file"],
        ":k8": ["@kubebuilder_linux//:file"],
    }),
    outs = [
        "kubebuilder",
        "etcd",
        "kube-apiserver",
    ],
    cmd = "cp $(SRCS) $(RULEDIR)",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

genrule(
    name = "fetch_kustomize",
    srcs = select({
        ":darwin": ["@kustomize_osx//:file"],
        ":k8": ["@kustomize_linux//:file"],
    }),
    outs = ["kustomize"],
    cmd = "cp $(SRCS) $@",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

genrule(
    name = "fetch_just",
    srcs = select({
        ":darwin": ["@just_osx//:file"],
        ":k8": ["@just_linux//:file"],
    }),
    outs = ["just"],
    cmd = "cp $(SRCS) $@",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

# boilerplate for handling linux vs OS versions of binary tools

config_setting(
    name = "k8",
    values = {"host_cpu": "k8"},
    visibility = ["//visibility:private"],
)

config_setting(
    name = "darwin",
    values = {"host_cpu": "darwin"},
    visibility = ["//visibility:private"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
