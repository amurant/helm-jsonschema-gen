package main

type ControllerDeployment struct {
	// Number of replicas
	ReplicaCount int `json:"replicaCount"`

	// Annotations to add to the deployment
	DeploymentAnnotations map[string]string `json:"deploymentAnnotations,omitempty"`

	// Annotations to add to the pods
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`
	// Annotations to add to the pods
	PodLabels map[string]string `json:"podLabels"`

	Strategy *DeploymentStrategy `json:"strategy,omitempty"`

	// Node labels for pod assignment
	NodeSelector map[string]string `json:"nodeSelector"`

	// Node affinity for pod assignment
	Affinity interface{} `json:"affinity"`

	// Node tolerations for pod assignment
	Tolerations []string `json:"tolerations"`

	Image Image `json:"image"`

	// CPU/memory resource requests/limits for the pods
	Resources interface{} `json:"resources"`

	// Optional additional arguments for controller
	ExtraArgs []string `json:"extraArgs"`

	// Optional additional environment variables for controller
	ExtraEnv []EnvVar `json:"extraEnv"`

	// Service account configuration for controller
	ServiceAccount ServiceAccount `json:"serviceAccount"`

	// Pod Security Context
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext interface{} `json:"securityContext"`

	// Container Security Context to be set on the controller component container
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext interface{} `json:"containerSecurityContext"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
