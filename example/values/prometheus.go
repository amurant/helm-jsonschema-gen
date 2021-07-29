package main

type Prometheus struct {
	// If `true`, enable Prometheus monitoring
	Enabled bool `json:"enabled"`

	Servicemonitor Servicemonitor `json:"servicemonitor"`
}

type Servicemonitor struct {
	// Enable Prometheus Operator ServiceMonitor monitoring
	Enabled bool `json:"enabled"`

	// Prometheus Instance definition
	PrometheusInstance string `json:"prometheusInstance"`

	//Prometheus scrape port
	TargetPort int `json:"targetPort"`

	// Prometheus scrape path
	Path string `json:"path"`

	// Prometheus scrape interval
	Interval string `json:"interval"`

	// Prometheus scrape timeout
	ScrapeTimeout string `json:"scrapeTimeout"`

	// Add custom labels to ServiceMonitor
	Labels map[string]string `json:"labels"`
}
