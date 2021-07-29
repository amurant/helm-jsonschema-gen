package main

type ServiceAccount struct {
	// If `true`, create a new service account for the cert-manager controller
	Create bool `json:"create"`

	// The name of the service account for the cert-manager controller to be used.
	// If not set and `serviceAccount.create` is `true`, a name is generated using
	// the fullname template
	Name string `json:"name,omitempty"`

	// Annotations to add to the service account for the cert-manager controller
	Annotations map[string]string `json:"annotations,omitempty"`

	// Automount API credentials for the cert-manager service account
	AutomountServiceAccountToken bool `json:"automountServiceAccountToken"`
}
