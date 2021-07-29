package main

type Probe struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	FailureThreshold int `json:"failureThreshold"`

	// Number of seconds after the container has started before liveness probes are initiated.
	// ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	InitialDelaySeconds int `json:"initialDelaySeconds"`

	// How often (in seconds) to perform the probe.
	PeriodSeconds int `json:"periodSeconds"`

	// Minimum consecutive successes for the probe to be considered successful after having failed.
	SuccessThreshold int `json:"successThreshold"`

	// Number of seconds after which the probe times out.
	// ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	TimeoutSeconds int `json:"timeoutSeconds"`
}
