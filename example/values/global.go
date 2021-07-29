package main

type Global struct {
	ImagePullSecrets  []interface{}     `json:"imagePullSecrets"`
	LeaderElection    LeaderElection    `json:"leaderElection"`
	LogLevel          int               `json:"logLevel"`
	PodSecurityPolicy PodSecurityPolicy `json:"podSecurityPolicy"`
	PriorityClassName string            `json:"priorityClassName"`
	Rbac              Rbac              `json:"rbac"`
}

type PodSecurityPolicy struct {
	// If `true`, create and use PodSecurityPolicy
	Enabled bool `json:"enabled"`

	// If `true`, use Apparmor seccomp profile in PSP
	UseAppArmor bool `json:"useAppArmor"`
}

type Rbac struct {
	Create bool `json:"create"`
}
