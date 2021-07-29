package main

import (
	"github.com/amurant/helm-jsonschema-gen/pkg/schema"
)

//go:generate go run github.com/amurant/helm-jsonschema-gen/cmd/helm-jsonschema-gen -p .

func main() {
	if err := schema.GenSchema(&defaultValues, "../values.schema.json"); err != nil {
		panic(err)
	}
	if err := schema.GenDefaults(&defaultValues, "../values.yaml"); err != nil {
		panic(err)
	}
	if err := schema.GenDocs(&defaultValues, "../values.docs.yaml"); err != nil {
		panic(err)
	}
}

var defaultValues = Values{
	Global: Global{
		ImagePullSecrets:  []interface{}{},
		PriorityClassName: "",
		LeaderElection: LeaderElection{
			Namespace: "kube-system",
		},
		LogLevel: 2,
		PodSecurityPolicy: PodSecurityPolicy{
			Enabled:     false,
			UseAppArmor: true,
		},
		Rbac: Rbac{
			Create: true,
		},
	},
	InstallCRDs: false,
	ControllerDeployment: ControllerDeployment{
		ReplicaCount: 1,
		Strategy:     DeploymentStrategy{},
		NodeSelector: map[string]string{},
		Affinity:     map[string]interface{}{},
		Tolerations:  []string{},
		Image: Image{
			Repository: "quay.io/jetstack/cert-manager-controller",
			PullPolicy: "IfNotPresent",
		},
		Resources: map[string]interface{}{},
		ExtraArgs: []string{},
		ExtraEnv:  []EnvVar{},
		ServiceAccount: ServiceAccount{
			Create:                       true,
			AutomountServiceAccountToken: true,
		},
		SecurityContext: map[string]interface{}{
			"runAsNonRoot": true,
		},
		ContainerSecurityContext: map[string]interface{}{},
	},
	Cainjector: Cainjector{
		Enabled: true,
		ControllerDeployment: ControllerDeployment{
			ReplicaCount: 1,
			Strategy:     DeploymentStrategy{},
			NodeSelector: map[string]string{},
			Affinity:     map[string]interface{}{},
			Tolerations:  []string{},
			Image: Image{
				Repository: "quay.io/jetstack/cert-manager-cainjector",
				PullPolicy: "IfNotPresent",
			},
			Resources: map[string]interface{}{},
			ExtraArgs: []string{},
			ExtraEnv:  []EnvVar{},
			ServiceAccount: ServiceAccount{
				Create:                       true,
				AutomountServiceAccountToken: true,
			},
			SecurityContext: map[string]interface{}{
				"runAsNonRoot": true,
			},
			ContainerSecurityContext: map[string]interface{}{},
		},
	},
	Prometheus: Prometheus{
		Enabled: false,
		Servicemonitor: Servicemonitor{
			PrometheusInstance: "default",
			TargetPort:         9402,
			Path:               "/metrics",
			Interval:           "60s",
			ScrapeTimeout:      "30s",
			Labels:             map[string]string{},
		},
	},
	Webhook: Webhook{
		ControllerDeployment: ControllerDeployment{
			ReplicaCount: 1,
			Strategy:     DeploymentStrategy{},
			NodeSelector: map[string]string{},
			Affinity:     map[string]interface{}{},
			Tolerations:  []string{},
			Image: Image{
				Repository: "quay.io/jetstack/cert-manager-webhook",
				PullPolicy: "IfNotPresent",
			},
			Resources: map[string]interface{}{},
			ExtraArgs: []string{},
			ExtraEnv:  []EnvVar{},
			ServiceAccount: ServiceAccount{
				Create:                       true,
				AutomountServiceAccountToken: true,
			},
			SecurityContext: map[string]interface{}{
				"runAsNonRoot": true,
			},
			ContainerSecurityContext: map[string]interface{}{},
		},
		LivenessProbe: Probe{
			FailureThreshold:    3,
			InitialDelaySeconds: 60,
			PeriodSeconds:       10,
			SuccessThreshold:    1,
			TimeoutSeconds:      1,
		},
		ReadinessProbe: Probe{
			FailureThreshold:    3,
			InitialDelaySeconds: 5,
			PeriodSeconds:       5,
			SuccessThreshold:    1,
			TimeoutSeconds:      1,
		},
		TimeoutSeconds:                          10,
		SecurePort:                              10250,
		HostNetwork:                             false,
		ServiceType:                             "ClusterIP",
		Url:                                     WebhookUrl{},
		MutatingWebhookConfigurationAnnotations: map[string]string{},
		ValidatingWebhookConfigurationAnnotations: map[string]string{},
	},
	ClusterResourceNamespace: "",
	FeatureGates:             "",
	ServiceLabels:            map[string]string{},
	VolumeMounts:             []interface{}{},
	Volumes:                  []interface{}{},
}
