GO = "@go_sdk//:bin/go"

CONTROLLER_GEN = "@io_k8s_sigs_controller_tools//cmd/controller-gen"

def gen_manifests(pkg_name, role_name = "manager-role", name = "gen_manifests"):
    """Generates manifests for a given controller-runtime project.

    pkg_name should be the bazel package name (e.g., scheduler)
    role_name is the name for the generated RBAC role
    """
    native.sh_binary(
      name = name,
      srcs = ["//hack/gen:gen-manifests.sh"],
      args = [
          "$(location %s)" % GO,
          "$(location %s)" % CONTROLLER_GEN,
          pkg_name,
          role_name
      ],
      data = [
          GO,
          CONTROLLER_GEN,
      ],
    )
