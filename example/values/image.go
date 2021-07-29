package main

type Image struct {
	// Image repository
	Repository string `json:"repository"`

	// You can manage a registry with
	// Example:
	//  registry: quay.io
	//  repository: jetstack/cert-manager-controller
	Registry string `json:"registry,omitempty"`

	// Override the image tag to deploy by setting this variable.
	// If no value is set, the chart's appVersion will be used.
	Tag string `json:"tag,omitempty,omitempty"`

	// Setting a digest will override any tag
	Digest string `json:"digest,omitempty"`

	// Image pull policy
	PullPolicy string `json:"pullPolicy,omitempty"`
}
