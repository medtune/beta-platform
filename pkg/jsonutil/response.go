package jsonutil

// SimpleResponse .
type SimpleResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Warnings []string    `json:"warnings"`
	Errors   []string    `json:"errors"`
}

// InferenceResult .
type InferenceResult struct {
	Keys   []string  `json:"keys"`
	Scores []float32 `json:"scores"`
}

// CamResult .
type CamResult struct {
	StaticPath string `json:"static_path"`
	URL        string `json:"url"`
}

// PackageVersion .
type PackageVersion struct {
	Major   string `json:"major"`
	Minor   string `json:"minor"`
	Patch   string `json:"patch"`
	Version string `json:"version"`
}

// ServiceStatus .
type ServiceStatus struct {
	Healthy  bool  `json:"healthy"`
	UnixTime int64 `json:"unix_time"`
}

// TestResponse .
type TestResponse struct {
	Test bool `json:"test"`
}
