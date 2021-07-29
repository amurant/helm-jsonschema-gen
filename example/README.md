| command | type | description | default |
| ------- | ---- | ----------- | ------- |
| --set installCRDs | boolean | If true, CRD resources will be installed as part of the Helm chart. If enabled, when uninstalling CRD resources will be deleted causing all installed custom resources to be DELETED. | `false` |
| --set replicaCount | integer | Number of replicas | `1` |
| --set deploymentAnnotations | map[string]string | Annotations to add to the deployment | `` |
| --set podAnnotations | map[string]string | Annotations to add to the pods | `` |
| --set podLabels | map[string]string | Annotations to add to the pods | `` |
| --set-string strategy.type | string | Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate. | `` |
| --set strategy.rollingUpdate.maxUnavailable | object | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods. | `` |
| --set strategy.rollingUpdate.maxSurge | object | The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods. | `` |
| --set nodeSelector | map[string]string | Node labels for pod assignment | `{}` |
| --set affinity | object | Node affinity for pod assignment | `{}` |
| --set tolerations | []string | Node tolerations for pod assignment | `[]` |
| --set-string image.repository | string | Image repository | `"quay.io/jetstack/cert-manager-controller"` |
| --set-string image.registry | string | You can manage a registry with Example:  registry: quay.io  repository: jetstack/cert-manager-controller | `` |
| --set-string image.tag | string | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. | `` |
| --set-string image.digest | string | Setting a digest will override any tag | `` |
| --set-string image.pullPolicy | string | Image pull policy | `"IfNotPresent"` |
| --set resources | object | CPU/memory resource requests/limits for the pods | `{}` |
| --set extraArgs | []string | Optional additional arguments for controller | `[]` |
| --set extraEnv | []object | Optional additional environment variables for controller | `[]` |
| --set serviceAccount.create | boolean | If `true`, create a new service account for the cert-manager controller | `true` |
| --set-string serviceAccount.name | string | The name of the service account for the cert-manager controller to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `` |
| --set serviceAccount.annotations | map[string]string | Annotations to add to the service account for the cert-manager controller | `` |
| --set serviceAccount.automountServiceAccountToken | boolean | Automount API credentials for the cert-manager service account | `true` |
| --set securityContext | object | Pod Security Context ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{"runAsNonRoot":true}` |
| --set containerSecurityContext | object | Container Security Context to be set on the controller component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| --set global.imagePullSecrets | []object | <nil> | `[]` |
| --set-string global.leaderElection.namespace | string | Override the namespace used to store the ConfigMap for leader election | `"kube-system"` |
| --set-string global.leaderElection.leaseDuration | string | The duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. | `` |
| --set-string global.leaderElection.renewDeadline | string | The interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. | `` |
| --set-string global.leaderElection.retryPeriod | string | The duration the clients should wait between attempting acquisition and renewal of a leadership. | `` |
| --set global.logLevel | integer | <nil> | `2` |
| --set global.podSecurityPolicy.enabled | boolean | If `true`, create and use PodSecurityPolicy | `false` |
| --set global.podSecurityPolicy.useAppArmor | boolean | If `true`, use Apparmor seccomp profile in PSP | `true` |
| --set-string global.priorityClassName | string | <nil> | `""` |
| --set global.rbac.create | boolean | <nil> | `true` |
| --set cainjector.enabled | boolean | <nil> | `true` |
| --set cainjector.replicaCount | integer | Number of replicas | `1` |
| --set cainjector.deploymentAnnotations | map[string]string | Annotations to add to the deployment | `` |
| --set cainjector.podAnnotations | map[string]string | Annotations to add to the pods | `` |
| --set cainjector.podLabels | map[string]string | Annotations to add to the pods | `` |
| --set-string cainjector.strategy.type | string | Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate. | `` |
| --set cainjector.strategy.rollingUpdate.maxUnavailable | object | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods. | `` |
| --set cainjector.strategy.rollingUpdate.maxSurge | object | The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods. | `` |
| --set cainjector.nodeSelector | map[string]string | Node labels for pod assignment | `{}` |
| --set cainjector.affinity | object | Node affinity for pod assignment | `{}` |
| --set cainjector.tolerations | []string | Node tolerations for pod assignment | `[]` |
| --set-string cainjector.image.repository | string | Image repository | `"quay.io/jetstack/cert-manager-cainjector"` |
| --set-string cainjector.image.registry | string | You can manage a registry with Example:  registry: quay.io  repository: jetstack/cert-manager-controller | `` |
| --set-string cainjector.image.tag | string | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. | `` |
| --set-string cainjector.image.digest | string | Setting a digest will override any tag | `` |
| --set-string cainjector.image.pullPolicy | string | Image pull policy | `"IfNotPresent"` |
| --set cainjector.resources | object | CPU/memory resource requests/limits for the pods | `{}` |
| --set cainjector.extraArgs | []string | Optional additional arguments for controller | `[]` |
| --set cainjector.extraEnv | []object | Optional additional environment variables for controller | `[]` |
| --set cainjector.serviceAccount.create | boolean | If `true`, create a new service account for the cert-manager controller | `true` |
| --set-string cainjector.serviceAccount.name | string | The name of the service account for the cert-manager controller to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `` |
| --set cainjector.serviceAccount.annotations | map[string]string | Annotations to add to the service account for the cert-manager controller | `` |
| --set cainjector.serviceAccount.automountServiceAccountToken | boolean | Automount API credentials for the cert-manager service account | `true` |
| --set cainjector.securityContext | object | Pod Security Context ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{"runAsNonRoot":true}` |
| --set cainjector.containerSecurityContext | object | Container Security Context to be set on the controller component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| --set prometheus.enabled | boolean | If `true`, enable Prometheus monitoring | `false` |
| --set prometheus.servicemonitor.enabled | boolean | Enable Prometheus Operator ServiceMonitor monitoring | `false` |
| --set-string prometheus.servicemonitor.prometheusInstance | string | Prometheus Instance definition | `"default"` |
| --set prometheus.servicemonitor.targetPort | integer | Prometheus scrape port | `9402` |
| --set-string prometheus.servicemonitor.path | string | Prometheus scrape path | `"/metrics"` |
| --set-string prometheus.servicemonitor.interval | string | Prometheus scrape interval | `"60s"` |
| --set-string prometheus.servicemonitor.scrapeTimeout | string | Prometheus scrape timeout | `"30s"` |
| --set prometheus.servicemonitor.labels | map[string]string | Add custom labels to ServiceMonitor | `{}` |
| --set webhook.replicaCount | integer | Number of replicas | `1` |
| --set webhook.deploymentAnnotations | map[string]string | Annotations to add to the deployment | `` |
| --set webhook.podAnnotations | map[string]string | Annotations to add to the pods | `` |
| --set webhook.podLabels | map[string]string | Annotations to add to the pods | `` |
| --set-string webhook.strategy.type | string | Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate. | `` |
| --set webhook.strategy.rollingUpdate.maxUnavailable | object | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods. | `` |
| --set webhook.strategy.rollingUpdate.maxSurge | object | The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods. | `` |
| --set webhook.nodeSelector | map[string]string | Node labels for pod assignment | `{}` |
| --set webhook.affinity | object | Node affinity for pod assignment | `{}` |
| --set webhook.tolerations | []string | Node tolerations for pod assignment | `[]` |
| --set-string webhook.image.repository | string | Image repository | `"quay.io/jetstack/cert-manager-webhook"` |
| --set-string webhook.image.registry | string | You can manage a registry with Example:  registry: quay.io  repository: jetstack/cert-manager-controller | `` |
| --set-string webhook.image.tag | string | Override the image tag to deploy by setting this variable. If no value is set, the chart's appVersion will be used. | `` |
| --set-string webhook.image.digest | string | Setting a digest will override any tag | `` |
| --set-string webhook.image.pullPolicy | string | Image pull policy | `"IfNotPresent"` |
| --set webhook.resources | object | CPU/memory resource requests/limits for the pods | `{}` |
| --set webhook.extraArgs | []string | Optional additional arguments for controller | `[]` |
| --set webhook.extraEnv | []object | Optional additional environment variables for controller | `[]` |
| --set webhook.serviceAccount.create | boolean | If `true`, create a new service account for the cert-manager controller | `true` |
| --set-string webhook.serviceAccount.name | string | The name of the service account for the cert-manager controller to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `` |
| --set webhook.serviceAccount.annotations | map[string]string | Annotations to add to the service account for the cert-manager controller | `` |
| --set webhook.serviceAccount.automountServiceAccountToken | boolean | Automount API credentials for the cert-manager service account | `true` |
| --set webhook.securityContext | object | Pod Security Context ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{"runAsNonRoot":true}` |
| --set webhook.containerSecurityContext | object | Container Security Context to be set on the controller component container ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ | `{}` |
| --set webhook.livenessProbe.failureThreshold | integer | Minimum consecutive failures for the probe to be considered failed after having succeeded. | `3` |
| --set webhook.livenessProbe.initialDelaySeconds | integer | Number of seconds after the container has started before liveness probes are initiated. ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | `60` |
| --set webhook.livenessProbe.periodSeconds | integer | How often (in seconds) to perform the probe. | `10` |
| --set webhook.livenessProbe.successThreshold | integer | Minimum consecutive successes for the probe to be considered successful after having failed. | `1` |
| --set webhook.livenessProbe.timeoutSeconds | integer | Number of seconds after which the probe times out. ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | `1` |
| --set webhook.readinessProbe.failureThreshold | integer | Minimum consecutive failures for the probe to be considered failed after having succeeded. | `3` |
| --set webhook.readinessProbe.initialDelaySeconds | integer | Number of seconds after the container has started before liveness probes are initiated. ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | `5` |
| --set webhook.readinessProbe.periodSeconds | integer | How often (in seconds) to perform the probe. | `5` |
| --set webhook.readinessProbe.successThreshold | integer | Minimum consecutive successes for the probe to be considered successful after having failed. | `1` |
| --set webhook.readinessProbe.timeoutSeconds | integer | Number of seconds after which the probe times out. ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | `1` |
| --set webhook.securePort | integer | The port that the webhook should listen on for requests. In GKE private clusters, by default kubernetes apiservers are allowed to talk to the cluster nodes only on 443 and 10250. so configuring securePort: 10250, will work out of the box without needing to add firewall rules or requiring NET_BIND_SERVICE capabilities to bind port numbers <1000 | `10250` |
| --set webhook.hostNetwork | boolean | Specifies if the webhook should be started in hostNetwork mode.  Required for use in some managed kubernetes clusters (such as AWS EKS) with custom CNI (such as calico), because control-plane managed by AWS cannot communicate with pods' IP CIDR and admission webhooks are not working  Since the default port for the webhook conflicts with kubelet on the host network, `webhook.securePort` should be changed to an available port if running in hostNetwork mode. | `false` |
| --set-string webhook.serviceType | string | The type of the `Service`. Specifies how the service should be handled. Useful if you want to expose the webhook to outside of the cluster. In some cases, the control plane cannot reach internal services. | `"ClusterIP"` |
| --set-string webhook.loadBalancerIP | string | The specific load balancer IP to use (when `serviceType` is `LoadBalancer`). | `` |
| --set webhook.timeoutSeconds | integer | Seconds the API server should wait the webhook to respond before treating the call as a failure. | `10` |
| --set-string webhook.url.host | string | <nil> | `` |
| --set webhook.mutatingWebhookConfigurationAnnotations | map[string]string | Annotations to add to the webhook MutatingWebhookConfiguration | `` |
| --set webhook.validatingWebhookConfigurationAnnotations | map[string]string | Annotations to add to the webhook ValidatingWebhookConfiguration | `` |
| --set-string clusterResourceNamespace | string | Override the namespace used to store DNS provider credentials etc. for ClusterIssuer resources. By default, the same namespace as cert-manager is deployed within is used. This namespace will not be automatically created by the Helm chart. | `""` |
| --set-string featureGates | string | Comma separated list of feature gates that should be enabled on the | `""` |
| --set-string http_proxy | string | Value of the `HTTP_PROXY` environment variable in the cert-manager pod | `` |
| --set-string https_proxy | string | Value of the `HTTPS_PROXY` environment variable in the cert-manager pod | `` |
| --set-string no_proxy | string | Value of the `NO_PROXY` environment variable in the cert-manager pod | `` |
| --set ingressShim | object | <nil> | `` |
| --set podDnsConfig | object | Optional cert-manager pod [DNS configurations](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-config) | `` |
| --set-string podDnsPolicy | string | Optional cert-manager pod [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-policy) Useful if you have a public and private DNS zone for the same domain on Route 53. What follows is an example of ensuring cert-manager can access an ingress or DNS TXT records at all times. **NOTE:** This requires Kubernetes 1.10 or `CustomPodDNS` feature gate enabled for the cluster to work. | `` |
| --set serviceLabels | map[string]string | Labels to add to the cert-manager controller service | `` |
| --set volumeMounts | []object | Volume mounts to add to cert-manager | `[]` |
| --set volumes | []object | Volumes to add to cert-manager | `[]` |
| --set-string creator | string | <nil> | `` |
