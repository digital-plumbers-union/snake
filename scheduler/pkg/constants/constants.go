package constants

const (
	// BuildNumberConfigMap is the ConfigMap where the scheduler will store the current build number
	// TODO: make configurable (see moodring?)
	BuildNumberConfigMap string = "snake-build-number"
	// BuildNumberLabel label representing the build number for a PipelineRun
	BuildNumberLabel string = "snake.dpu.sh/build-number"
	// AssignBuildNumberAnnotation annotation should be added to Pipelines which want their PipelineRuns to
	// be assigned a build number
	AssignBuildNumberAnnotation string = "snake.dpu.sh/assign-build-number"
)
