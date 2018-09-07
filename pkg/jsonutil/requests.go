package jsonutil

// LoginData .
type LoginData struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// SignupData .
type SignupData struct {
	Secret          string `json:"secret" valid:"required"`
	Email           string `json:"email" valid:"email"`
	Username        string `json:"username" valid:"required"`
	Password        string `json:"password" valid:"required"`
	PasswordConfirm string `json:"passwordc" valid:"required"`
}

// RunImageInference .
type RunImageInference struct {
	Id       int    `json:"id"`
	Image    string `json:"image"`
	File     string `json:"file"`
	NumPreds string `json:"numpreds"`
	ModelID  string `json:"model_id"`
}

// RunImageCam .
type RunImageCam struct {
	Id      int    `json:"id"`
	Target  string `json:"target" valid:"required"`
	Output  string `json:"output"`
	Force   bool   `json:"force"`
	ModelID string `json:"model_id" valid:"required"`
}

// GetStatus .
type GetStatus struct {
	ModelID string `json:"model_id" valid:"required"`
}

// DataDrop .
type DataDrop struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
	File  string `json:"file"`
}

// ProcessImage .
type ProcessImage struct {
	Id       int    `json:"id"`
	Target   string `json:"target"`
	NumPreds string `json:"numpreds"`
	Output   string `json:"output"`
	ModelID  string `json:"model_id"`
}

// GetPathologyAL .
type GetPathologyAL struct {
	Pathology string `json:"pathology" valid:"required"`
}

// GetSpecPoolGrid .
type GetSpecPoolGrid struct {
	Ignore []string `json:"ignore"`
}

// TestRequest .
type TestRequest struct {
	Test  bool `json:"test"`
	Sleep int  `json:"sleep"`
}

// Job .
type Job struct {
	Name            string            `json:"name"`
	Requires        string            `json:"requires"`
	Function        string            `json:"function"`
	InferenceParams RunImageInference `json:"inference_params"`
	CamParams       RunImageCam       `json:"cam_params"`
	Context         string            `json:"context"`
}

// CustomExecutionRequest .
type CustomExecutionRequest struct {
	Id            string `json:"id" valid:"required"`
	Name          string `json:"name" valid:"required"`
	Jobs          []*Job `json:"jobs"`
	Concurrency   bool   `json:"concurrency"`
	MaxGoroutines int    `json:"max_goroutines"`
	Context       string `json:"context" valid:"required"`
}

/*
# Request example
{
	"id" : "J9E83U2J9",
	"name" : "client-job-num-1"
	"concurrency" : true,
	"max_goroutines": 8,
	"context" : "timeout 5",
	"jobs" : [
		{
			"name" : "simple-inference",
			"requires" : "",
			"function" : "demo/mura/run_inference",
			"inference_params" : {
				"file" : "image_0.png"
			},
			"cam_params" : {},
		},
		{
			"name" : "simple-inference",
			"requires" : "",
			"type": "cam",
			"function" : "demo/mura/run_cam",
			"inference_params" : {},
			"cam_params" : {
				"file": "image_0.png",
			},
		},
	]
}

# Response example
{
	"id" : "J9E83U2J9",
	"name" : "client-job-num-1",
	"round_trip" : 1.392892,
	"execution_round_trip" : 1.1232,
	"success" : true,
	"warning" : [],
	"errors" : [],
	"jobs" : [
		{
			"name" : "simple-inference",
			"round_trip" : 1.008992,
			"success" : true,
			"errors": [],
			"warnings": [],
		},
		{
			"name" : "simple-inference",
			"round_trip" : 0.9598,
			"success" : true,
			"errors": [],
			"warnings": [],
		},
	],
}
*/
