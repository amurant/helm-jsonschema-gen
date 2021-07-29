package main

type WebhookUrl struct {
	Host string `json:"host,omitempty"`
}

type Webhook struct {
	ControllerDeployment

	LivenessProbe  Probe `json:"livenessProbe"`
	ReadinessProbe Probe `json:"readinessProbe"`

	// The port that the webhook should listen on for requests.
	// In GKE private clusters, by default kubernetes apiservers are allowed to
	// talk to the cluster nodes only on 443 and 10250. so configuring
	// securePort: 10250, will work out of the box without needing to add firewall
	// rules or requiring NET_BIND_SERVICE capabilities to bind port numbers <1000
	SecurePort int `json:"securePort"`

	// Specifies if the webhook should be started in hostNetwork mode.
	//
	// Required for use in some managed kubernetes clusters (such as AWS EKS) with custom
	// CNI (such as calico), because control-plane managed by AWS cannot communicate
	// with pods' IP CIDR and admission webhooks are not working
	//
	// Since the default port for the webhook conflicts with kubelet on the host
	// network, `webhook.securePort` should be changed to an available port if
	// running in hostNetwork mode.
	HostNetwork bool `json:"hostNetwork"`

	// The type of the `Service`.
	// Specifies how the service should be handled. Useful if you want to expose the
	// webhook to outside of the cluster. In some cases, the control plane cannot
	// reach internal services.
	ServiceType string `json:"serviceType"`

	// The specific load balancer IP to use (when `serviceType` is `LoadBalancer`).
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`

	// Seconds the API server should wait the webhook to respond before treating the call as a failure.
	TimeoutSeconds int `json:"timeoutSeconds"`

	// Overrides the mutating webhook and validating webhook so they reach the webhook
	// service using the `url` field instead of a service.
	Url WebhookUrl `json:"url"`

	// Annotations to add to the webhook MutatingWebhookConfiguration
	MutatingWebhookConfigurationAnnotations map[string]string `json:"mutatingWebhookConfigurationAnnotations,omitempty"`

	// Annotations to add to the webhook ValidatingWebhookConfiguration
	ValidatingWebhookConfigurationAnnotations map[string]string `json:"validatingWebhookConfigurationAnnotations,omitempty"`
}
