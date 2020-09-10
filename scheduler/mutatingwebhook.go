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

// +kubebuilder:webhook:path=/mutate-v1beta1-pipelinerun,mutating=true,failurePolicy=fail,groups=tekton.dev,resources=pipelineruns,verbs=create,versions=v1beta1,name=scheduler.snake.dpu.sh,sideEffects=some

// pipelineRunAnnotator annotates Pods
type pipelineRunAnnotator struct {
	Client  client.Client
	decoder *admission.Decoder
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
	// TODO: make sure labels exist (?)
	// add the incremented build number as a label
	run.Labels[constants.BuildNumberLabel] = strconv.Itoa(a.BuildNumber)
	// extract build from configMap and annotate it
	buildNumberConfigMap := &corev1.ConfigMap{}
	if err := a.Client.Get(ctx, types.NamespacedName{Name: constants.BuildNumberConfigMap, Namespace: run.ObjectMeta.Namespace}, buildNumberConfigMap); err != nil {
		// configMap not found error
		return admission.Errored(1, err)
	}

	buildNumber := buildNumberConfigMap.Data[constants.BuildNumberKey]
	// update the annotation
	run.Annotations["snake.blockheads.info/build-number"] = buildNumber
	// update the env variables

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
