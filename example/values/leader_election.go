package main

type LeaderElection struct {
	// Override the namespace used to store the ConfigMap for leader election
	Namespace string `json:"namespace"`

	// The duration that non-leader candidates will wait after observing a
	// leadership renewal until attempting to acquire leadership of a led but
	// unrenewed leader slot. This is effectively the maximum duration that a
	// leader can be stopped before it is replaced by another candidate.
	LeaseDuration string `json:"leaseDuration,omitempty"`

	// The interval between attempts by the acting master to renew a leadership
	// slot before it stops leading. This must be less than or equal to the
	// lease duration.
	RenewDeadline string `json:"renewDeadline,omitempty"`

	// The duration the clients should wait between attempting acquisition and
	// renewal of a leadership.
	RetryPeriod string `json:"retryPeriod,omitempty"`
}
