/*
Header boilerplate comment
*/
/* DO NOT EDIT, this file was auto-generated */

package main

func (ControllerDeployment) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Affinity":
		return `Node affinity for pod assignment`
	case "ContainerSecurityContext":
		return "Container Security Context to be set on the controller component container\nref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/"
	case "DeploymentAnnotations":
		return `Annotations to add to the deployment`
	case "ExtraArgs":
		return `Optional additional arguments for controller`
	case "ExtraEnv":
		return `Optional additional environment variables for controller`
	case "NodeSelector":
		return `Node labels for pod assignment`
	case "PodAnnotations":
		return `Annotations to add to the pods`
	case "PodLabels":
		return `Annotations to add to the pods`
	case "ReplicaCount":
		return `Number of replicas`
	case "Resources":
		return `CPU/memory resource requests/limits for the pods`
	case "SecurityContext":
		return "Pod Security Context\nref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/"
	case "ServiceAccount":
		return `Service account configuration for controller`
	case "Tolerations":
		return `Node tolerations for pod assignment`
	default:
		return ""
	}
}
func (DeploymentStrategy) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "RollingUpdate":
		return `Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.`
	case "Type":
		return `Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.`
	default:
		return ""
	}
}
func (Image) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Digest":
		return `Setting a digest will override any tag`
	case "PullPolicy":
		return `Image pull policy`
	case "Registry":
		return "You can manage a registry with\nExample:\n registry: quay.io\n repository: jetstack/cert-manager-controller"
	case "Repository":
		return `Image repository`
	case "Tag":
		return "Override the image tag to deploy by setting this variable.\nIf no value is set, the chart's appVersion will be used."
	default:
		return ""
	}
}
func (LeaderElection) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "LeaseDuration":
		return "The duration that non-leader candidates will wait after observing a\nleadership renewal until attempting to acquire leadership of a led but\nunrenewed leader slot. This is effectively the maximum duration that a\nleader can be stopped before it is replaced by another candidate."
	case "Namespace":
		return `Override the namespace used to store the ConfigMap for leader election`
	case "RenewDeadline":
		return "The interval between attempts by the acting master to renew a leadership\nslot before it stops leading. This must be less than or equal to the\nlease duration."
	case "RetryPeriod":
		return "The duration the clients should wait between attempting acquisition and\nrenewal of a leadership."
	default:
		return ""
	}
}
func (PodSecurityPolicy) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Enabled":
		return "If `true`, create and use PodSecurityPolicy"
	case "UseAppArmor":
		return "If `true`, use Apparmor seccomp profile in PSP"
	default:
		return ""
	}
}
func (Probe) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "FailureThreshold":
		return `Minimum consecutive failures for the probe to be considered failed after having succeeded.`
	case "InitialDelaySeconds":
		return "Number of seconds after the container has started before liveness probes are initiated.\nref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes"
	case "PeriodSeconds":
		return `How often (in seconds) to perform the probe.`
	case "SuccessThreshold":
		return `Minimum consecutive successes for the probe to be considered successful after having failed.`
	case "TimeoutSeconds":
		return "Number of seconds after which the probe times out.\nref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes"
	default:
		return ""
	}
}
func (Prometheus) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Enabled":
		return "If `true`, enable Prometheus monitoring"
	default:
		return ""
	}
}
func (RollingUpdateDeployment) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "MaxSurge":
		return "The maximum number of pods that can be scheduled above the desired number of\npods.\nValue can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).\nThis can not be 0 if MaxUnavailable is 0.\nAbsolute number is calculated from percentage by rounding up.\nDefaults to 25%.\nExample: when this is set to 30%, the new ReplicaSet can be scaled up immediately when\nthe rolling update starts, such that the total number of old and new pods do not exceed\n130% of desired pods. Once old pods have been killed,\nnew ReplicaSet can be scaled up further, ensuring that total number of pods running\nat any time during the update is at most 130% of desired pods."
	case "MaxUnavailable":
		return "The maximum number of pods that can be unavailable during the update.\nValue can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).\nAbsolute number is calculated from percentage by rounding down.\nThis can not be 0 if MaxSurge is 0.\nDefaults to 25%.\nExample: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods\nimmediately when the rolling update starts. Once new pods are ready, old ReplicaSet\ncan be scaled down further, followed by scaling up the new ReplicaSet, ensuring\nthat the total number of pods available at all times during the update is at\nleast 70% of desired pods."
	default:
		return ""
	}
}
func (ServiceAccount) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Annotations":
		return `Annotations to add to the service account for the cert-manager controller`
	case "AutomountServiceAccountToken":
		return `Automount API credentials for the cert-manager service account`
	case "Create":
		return "If `true`, create a new service account for the cert-manager controller"
	case "Name":
		return "The name of the service account for the cert-manager controller to be used.\nIf not set and `serviceAccount.create` is `true`, a name is generated using\nthe fullname template"
	default:
		return ""
	}
}
func (Servicemonitor) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "Enabled":
		return `Enable Prometheus Operator ServiceMonitor monitoring`
	case "Interval":
		return `Prometheus scrape interval`
	case "Labels":
		return `Add custom labels to ServiceMonitor`
	case "Path":
		return `Prometheus scrape path`
	case "PrometheusInstance":
		return `Prometheus Instance definition`
	case "ScrapeTimeout":
		return `Prometheus scrape timeout`
	case "TargetPort":
		return `Prometheus scrape port`
	default:
		return ""
	}
}
func (Values) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "ClusterResourceNamespace":
		return "Override the namespace used to store DNS provider credentials etc. for ClusterIssuer\nresources. By default, the same namespace as cert-manager is deployed within is\nused. This namespace will not be automatically created by the Helm chart."
	case "FeatureGates":
		return `Comma separated list of feature gates that should be enabled on the`
	case "HttpProxy":
		return "Value of the `HTTP_PROXY` environment variable in the cert-manager pod"
	case "HttpsProxy":
		return "Value of the `HTTPS_PROXY` environment variable in the cert-manager pod"
	case "InstallCRDs":
		return "If true, CRD resources will be installed as part of the Helm chart.\nIf enabled, when uninstalling CRD resources will be deleted causing all\ninstalled custom resources to be DELETED."
	case "NoProxy":
		return "Value of the `NO_PROXY` environment variable in the cert-manager pod"
	case "PodDnsConfig":
		return `Optional cert-manager pod [DNS configurations](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-config)`
	case "PodDnsPolicy":
		return "Optional cert-manager pod [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-policy)\nUseful if you have a public and private DNS zone for\nthe same domain on Route 53. What follows is an example of ensuring\ncert-manager can access an ingress or DNS TXT records at all times.\n**NOTE:** This requires Kubernetes 1.10 or `CustomPodDNS` feature\ngate enabled for the cluster to work."
	case "ServiceLabels":
		return `Labels to add to the cert-manager controller service`
	case "VolumeMounts":
		return `Volume mounts to add to cert-manager`
	case "Volumes":
		return `Volumes to add to cert-manager`
	default:
		return ""
	}
}
func (Webhook) GetFieldDocString(fieldName string) string {
	switch fieldName {
	case "HostNetwork":
		return "Specifies if the webhook should be started in hostNetwork mode.\n\nRequired for use in some managed kubernetes clusters (such as AWS EKS) with custom\nCNI (such as calico), because control-plane managed by AWS cannot communicate\nwith pods' IP CIDR and admission webhooks are not working\n\nSince the default port for the webhook conflicts with kubelet on the host\nnetwork, `webhook.securePort` should be changed to an available port if\nrunning in hostNetwork mode."
	case "LoadBalancerIP":
		return "The specific load balancer IP to use (when `serviceType` is `LoadBalancer`)."
	case "MutatingWebhookConfigurationAnnotations":
		return `Annotations to add to the webhook MutatingWebhookConfiguration`
	case "SecurePort":
		return "The port that the webhook should listen on for requests.\nIn GKE private clusters, by default kubernetes apiservers are allowed to\ntalk to the cluster nodes only on 443 and 10250. so configuring\nsecurePort: 10250, will work out of the box without needing to add firewall\nrules or requiring NET_BIND_SERVICE capabilities to bind port numbers <1000"
	case "ServiceType":
		return "The type of the `Service`.\nSpecifies how the service should be handled. Useful if you want to expose the\nwebhook to outside of the cluster. In some cases, the control plane cannot\nreach internal services."
	case "TimeoutSeconds":
		return `Seconds the API server should wait the webhook to respond before treating the call as a failure.`
	case "Url":
		return "Overrides the mutating webhook and validating webhook so they reach the webhook\nservice using the `url` field instead of a service."
	case "ValidatingWebhookConfigurationAnnotations":
		return `Annotations to add to the webhook ValidatingWebhookConfiguration`
	default:
		return ""
	}
}
