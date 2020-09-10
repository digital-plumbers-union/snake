package main

import (
	"context"
	"os"
	"strconv"

	constants "github.com/digital-plumbers-union/snake/scheduler/pkg/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;create;update

var (
	buildNumber int
	namespace   string
)

func init() {
	log.SetLogger(zap.New())
	namespace = os.Getenv("POD_NAMESPACE")
}

func main() {
	entryLog := log.Log.WithName("entrypoint")

	// Setup a Manager
	entryLog.Info("setting up manager")
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		entryLog.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	buildNumber, err = initializeBuildNumber(mgr.GetClient(), namespace)
	if err != nil {
		entryLog.Error(err, "Failed to initialize the build number")
		os.Exit(1)
	}
	if buildNumber == 0 {
		entryLog.Info("Could not find the build number ConfigMap. Created the ConfigMap instead")
	}
	entryLog.Info("Initialized build number: ", buildNumber)

	// Setup webhooks
	entryLog.Info("setting up webhook server")
	hookServer := mgr.GetWebhookServer()

	entryLog.Info("registering webhooks to the webhook server")
	hookServer.Register("/mutate-v1beta1-pipelinerun", &webhook.Admission{Handler: &pipelineRunAnnotator{Client: mgr.GetClient(), Namespace: namespace, BuildNumber: buildNumber}})

	entryLog.Info("starting manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		entryLog.Error(err, "unable to run manager")
		os.Exit(1)
	}
}

func initializeBuildNumber(client client.Client, namespace string) (int, error) {
	// Fetch the ConfigMap from the cache
	cm := &corev1.ConfigMap{}
	err := client.Get(context.Background(), types.NamespacedName{Name: constants.BuildNumberConfigMap, Namespace: namespace}, cm)

	if err != nil {
		// configMap not found, creating it
		configData := make(map[string]string)
		configData[constants.BuildNumberKey] = "0"
		newCm := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      constants.BuildNumberConfigMap,
				Namespace: namespace,
			},
			Data: configData,
		}
		if err := client.Create(context.Background(), newCm); err != nil {
			return 0, err
		}
		return strconv.Atoi(configData[constants.BuildNumberKey])
	}

	// otherwise we found the configmap, attempt to parse the build number
	return strconv.Atoi(cm.Data[constants.BuildNumberKey])
}
