package main

type Values struct {
	// If true, CRD resources will be installed as part of the Helm chart.
	// If enabled, when uninstalling CRD resources will be deleted causing all
	// installed custom resources to be DELETED.
	InstallCRDs bool `json:"installCRDs"`

	ControllerDeployment

	Global     Global     `json:"global"`
	Cainjector Cainjector `json:"cainjector"`
	Prometheus Prometheus `json:"prometheus"`
	Webhook    Webhook    `json:"webhook"`

	// Override the namespace used to store DNS provider credentials etc. for ClusterIssuer
	// resources. By default, the same namespace as cert-manager is deployed within is
	// used. This namespace will not be automatically created by the Helm chart.
	ClusterResourceNamespace string `json:"clusterResourceNamespace"`

	//Comma separated list of feature gates that should be enabled on the
	FeatureGates string `json:"featureGates"`

	// Value of the `HTTP_PROXY` environment variable in the cert-manager pod
	HttpProxy string `json:"http_proxy,omitempty"`
	// Value of the `HTTPS_PROXY` environment variable in the cert-manager pod
	HttpsProxy string `json:"https_proxy,omitempty"`
	// Value of the `NO_PROXY` environment variable in the cert-manager pod
	NoProxy string `json:"no_proxy,omitempty"`

	IngressShim interface{} `json:"ingressShim"`

	// Optional cert-manager pod [DNS configurations](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-config)
	PodDnsConfig interface{} `json:"podDnsConfig,omitempty"`

	// Optional cert-manager pod [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-policy)
	// Useful if you have a public and private DNS zone for
	// the same domain on Route 53. What follows is an example of ensuring
	// cert-manager can access an ingress or DNS TXT records at all times.
	// **NOTE:** This requires Kubernetes 1.10 or `CustomPodDNS` feature
	// gate enabled for the cluster to work.
	PodDnsPolicy string `json:"podDnsPolicy,omitempty"`

	// Labels to add to the cert-manager controller service
	ServiceLabels map[string]string `json:"serviceLabels,omitempty"`

	// Volume mounts to add to cert-manager
	VolumeMounts []interface{} `json:"volumeMounts"`

	// Volumes to add to cert-manager
	Volumes []interface{} `json:"volumes"`

	Creator string `json:"creator,omitempty" jsonschema:"enum=static,enum=helm"`
}
