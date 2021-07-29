package main

// +k8s:openapi-gen=true
type Cainjector struct {
	Enabled bool `json:"enabled"`

	ControllerDeployment
}
