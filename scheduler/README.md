# Snake Scheduler

Tekton resource MutatingAdmissionWebhook.

## Functionality

### Incremental Build Number Generation

#### Tracking the Build Number

##### Startup

On startup, the controller should look for a ConfigMap (`snake-build-number`) in the namespace that the pod is deployed in.  If the ConfigMap is found, the value for the only key in the ConfigMap should be converted to an int and used to initialize the build number passed to the mutating webhook.  If the ConfigMap is not found, the build number should be initialized to 0, and the ConfigMap created and initialized to 0.

##### Recording

`snake-build-number` should be updated by the mutating webhook every few seconds as well as updated when the controller is terminated, to minimize the amount of duplicate builds ids in the event of a crash or rescheduling.

#### Mutating Webhook Logic

If a PipelineRun is scheduled with the `snake.blockheads.info/generate-build-number=true` annotation, then this controller:

- Increments the build number stored in memory
- Adds it to the PipelineRun as the value for the `snake.blockheads.info/build-number` label
- Adds (or modifies) a `podTemplate` to the PipelineRun ([see the Tekton documentation](https://github.com/tektoncd/pipeline/blob/master/docs/pipelineruns.md#specifying-a-pod-template)) which sets the environment variables:
  - `BUILD_NUMBER`, `JFROG_CLI_BUILD_NUMBER` to the incremented build number
  - `JFROG_CLI_BUILD_NAME` to the defined constant for our Jfrog Artifactory build name that is used for all builds (see `hack/constants/constants.json` and `hack/constants/constants.go` for the definition of those constants, see `go/snake` for usage of those constants)


