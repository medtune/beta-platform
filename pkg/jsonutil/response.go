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

// Success return a success json response
// equavalent to {"success" : true}
func Success() *SimpleResponse {
	return &SimpleResponse{
		Success: true,
	}
}

// SuccessData return a success json response with
// extra data informations
// example: {"success" : true, "data": {"name": "r1",
// "test": true}}
func SuccessData(data interface{}) *SimpleResponse {
	return &SimpleResponse{
		Success: true,
		Data:    data,
	}
}

// Fail return a non success reponse with success = false
// adding all passed errors to the response
// example {"success" : false, "errors": ["can't find my self",
// "where am i "]}
func Fail(errors ...error) *SimpleResponse {
	errs := make([]string, 0, len(errors))
	for _, err := range errors {
		errs = append(errs, err.Error())
	}
	return &SimpleResponse{
		Success: false,
		Errors:  errs,
	}
}
