package(default_visibility = ["//visibility:public"])

GO = "@go_sdk//:bin/go"

CONTROLLER_GEN = "@io_k8s_sigs_controller_tools//cmd/controller-gen"

sh_binary(
    name = "gen-manifests",
    srcs = ["gen-manifests.sh"],
    args = [
        "$(location %s)" % GO,
        "$(location %s)" % CONTROLLER_GEN,
    ],
    data = [
        GO,
        CONTROLLER_GEN,
    ],
)

sh_binary(
    name = "gen-code",
    srcs = ["gen-code.sh"],
    args = [
        "$(location %s)" % GO,
        "$(location %s)" % CONTROLLER_GEN,
    ],
    data = [
        GO,
        CONTROLLER_GEN,
        "boilerplate.go.txt",
    ],
)
