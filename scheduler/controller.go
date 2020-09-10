package main

import (
	"context"
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

//  reconcileConfigMap reconciles ConfigMap
type reconcileConfigMap struct {
	// client can be used to retrieve objects from the APIServer.
	client client.Client
}

// Implement reconcile.Reconciler so the controller can reconcile objects
var _ reconcile.Reconciler = &reconcileConfigMap{}

func (r *reconcileConfigMap) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	// set up a convenient log object so we don't have to type request over and over again
	log := log.FromContext(ctx)

	// Fetch the ConfigMap from the cache
	cm := &corev1.ConfigMap{}
	err := r.client.Get(ctx, types.NamespacedName{Name: "snake-build-number", Namespace: request.Namespace}, cm)

	if err != nil {
		log.Info("Could not find ConfigMap snake-build-number. Creating the ConfigMap instead")
		// configMap not found, creating it
		configData := make(map[string]string)
		configData["buildNumber"] = "0"
		newCm := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "snake-build-number",
				Namespace: request.Namespace,
			},
			Data: configData,
		}
		if err := r.client.Create(context.TODO(), newCm); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	// Print the ConfigMap
	log.Info("Reconciling ConfigMap", "ConfigMap name", cm.ObjectMeta.Name)

	// Update buildNumber
	buildNumber, _ := strconv.Atoi(cm.Data["buildNumber"])
	buildNumber = buildNumber + 1

	// Update ConfigMap
	cm.Data["buildNumber"] = strconv.Itoa(buildNumber)
	if err := r.client.Update(context.TODO(), cm); err != nil {
		return reconcile.Result{}, fmt.Errorf("could not update ConfigMap: %+v", err)
	}

	return reconcile.Result{}, nil
}
