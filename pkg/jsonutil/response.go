package jsonutil

import (
	"time"
)

// DefaultResponse .
type DefaultResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data,omitempty"`
	Warnings []string    `json:"warnings,omitempty"`
	Errors   []string    `json:"errors,omitempty"`
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
	URL       string        `json:"url,omitempty"`
	ModelID   string        `json:"model_id"`
	RoundTrip time.Duration `json:"round_trip"`
}

// StatusResult .
type StatusResult struct {
	ModelID   string        `json:"model_id"`
	Status    string        `json:"status"`
	Version   int64         `json:"version"`
	RoundTrip time.Duration `json:"round_trip"`
}

// ProcessResult .
type ProcessResult struct {
	Inference *InferenceResult `json:"inference,omitempty"`
	Cam       *CamResult       `json:"cam,omitempty"`
	ModelID   string           `json:"model_id"`
	Timing    time.Duration    `json:"timing"`
}

// ServiceStatus .
type ServiceStatus struct {
	Healthy   bool          `json:"healthy"`
	UnixTime  int64         `json:"unix_time"`
	RoundTrip time.Duration `json:"round_trip"`
}

// TestResponse .
type TestResponse struct {
	Test bool `json:"test"`
}

// GetPathologyALResponse .
type GetPathologyALResponse struct {
	Pathology string      `json:"pathology"`
	Table     interface{} `json:"table"`
}

// GetSpecPoolGridResponse .
type GetSpecPoolGridResponse struct {
	Specs interface{} `json:"specs"`
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

type CapsulHealthCheckResponse struct {
	Healthy   bool          `json:"healthy"`
	Status    string        `json:"status"`
	Version   int64         `json:"version"`
	RoundTrip time.Duration `json:"round_trip"`
}

// CapsulGlobalHealthCheckResponse .
type CapsulGlobalHealthCheckResponse struct {
	RoundTrip time.Duration                `json:"round_trip"`
	Report    []*CapsulHealthCheckResponse `json:"report"`
}
