# `snake`

Helps unclog Tekton Pipelines.

>A plumber's snake or drain snake is a slender, flexible auger used to dislodge clogs in plumbing. The plumber's snake is often reserved for difficult clogs that cannot be loosened with a plunger. It is also sometimes called a toilet jack. A plumbers snake is often used by plumbers to clear a clogged drain pipe or sanitary sewer. ([Wikipedia](https://en.wikipedia.org/wiki/Plumber%27s_snake)).

`snake` is a pile of loosely related Go binaries that is related to making sure pipelines move smoothly.

- `pkg` library code that is not specific to any individual `snake` binary.
- `scheduler` is a Tekton MutatingAdmissionWebhook that customizes scheduled Tekton resources (e.g, adding incremental build numbers)


## Development

### Requirements

- [Bazelisk](https://github.com/bazelbuild/bazelisk)

>It is recommended to install [`just@v0.5.8`](https://github.com/casey/just/releases/tag/v0.5.8), but users without `just` installed can simply run `bin/just` instead after running `hack/tools/install.sh`.
