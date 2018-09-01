package jsonutil

import "time"

// DefaultResponse .
type DefaultResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Warnings []string    `json:"warnings"`
	Errors   []string    `json:"errors"`
}

// InferenceResult .
type InferenceResult struct {
	Keys      []string      `json:"keys"`
	Scores    []float32     `json:"scores"`
	ModelID   string        `json:"model_id"`
	RoundTrip time.Duration `json:"round_trip"`
}

// CamResult .
type CamResult struct {
	Output    string        `json:"output"`
	URL       string        `json:"url"`
	ModelID   string        `json:"model_id"`
	RoundTrip time.Duration `json:"round_trip"`
}

type StatusResult struct {
	Status    string        `json:"status"`
	Version   int64         `json:"version"`
	RoundTrip time.Duration `json:"round_trip"`
}

// ProcessResult .
type ProcessResult struct {
	Inference *InferenceResult `json:"inference"`
	Cam       *CamResult       `json:"cam"`
	ModelID   string           `json:"model_id"`
	Timing    time.Duration    `json:"timing"`
}

// PackageVersion .
type PackageVersion struct {
	Major   int8   `json:"major"`
	Minor   int8   `json:"minor"`
	Patch   int8   `json:"patch"`
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

// JobResult .
type JobResult struct {
	Name            string          `json:"name"`
	Success         bool            `json:"success"`
	Errors          []string        `json:"errors"`
	Warnings        []string        `json:"warnings"`
	RoundTrip       time.Duration   `json:"round_trip"`
	CamResult       CamResult       `json:"cam_result"`
	InferenceResult InferenceResult `json:"inference_result"`
}

// CustomExecutionResponse .
type CustomExecutionResponse struct {
	Id                 string        `json:"id"`
	Name               string        `json:"name"`
	Success            bool          `json:"success"`
	RoundTrip          time.Duration `json:"round_trip"`
	ExecutionRoundTrip time.Duration `json:"execution_round_trip"`
	Jobs               []*JobResult  `json:"jobs"`
	Errors             []string      `json:"errors"`
	Warnings           []string      `json:"warnings"`
}
