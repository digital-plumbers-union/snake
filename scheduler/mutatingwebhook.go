package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	constants "github.com/digital-plumbers-union/snake/scheduler/pkg/constants"
	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-v1beta1-pipelinerun,mutating=true,failurePolicy=fail,groups=tekton.dev,resources=pipelineruns,verbs=create,versions=v1beta1,name=scheduler.snake.dpu.sh,sideEffects=Some

// pipelineRunAnnotator annotates PipelineRuns
type pipelineRunAnnotator struct {
	Client      client.Client
	decoder     *admission.Decoder
	BuildNumber int
	// the namespace the server is deployed into, used to access/update the configmap
	Namespace string
}

// pipelineRunAnnotator generates a build number for pipelineruns with a specific annotation
func (a *pipelineRunAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	run := &pipelinev1.PipelineRun{}

	err := a.decoder.Decode(req, run)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// return Allowed early if we dont find the annotation telling us to generate a
	// build number
	if !(metav1.HasAnnotation(run.ObjectMeta, constants.AssignBuildNumberAnnotation)) {
		return admission.Allowed(fmt.Sprintf("Annotation %s not found", constants.AssignBuildNumberAnnotation))
	}

	// increment build, convert to string
	a.BuildNumber = a.BuildNumber + 1
	strBuildNumber := strconv.Itoa(a.BuildNumber)
	// TODO: make sure labels object exists on our PRun (?)
	// add the incremented build number as a label to our PipelineRun
	run.Labels[constants.BuildNumberLabel] = strBuildNumber

	// update the configmap storing the build number
	// get configmap
	buildNumberConfigMap := &corev1.ConfigMap{}
	if err := a.Client.Get(ctx, types.NamespacedName{Name: constants.BuildNumberConfigMap, Namespace: a.Namespace}, buildNumberConfigMap); err != nil {
		// configMap not found error
		return admission.Errored(1, err)
	}

	// update the key storing the build number
	buildNumberConfigMap.Data[constants.BuildNumberKey] = strBuildNumber

	// update the configmap
	if err := a.Client.Update(ctx, buildNumberConfigMap); err != nil {
		// failed to update configMap error
		return admission.Errored(1, err)
	}

	// marshal our modified PipelineRun so we can return the patch
	marshaledRun, err := json.Marshal(run)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledRun)
}

// pipelineRunAnnotator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (a *pipelineRunAnnotator) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}
