package jsonutil

type SimpleResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Warnings []string    `json:"warnings"`
	Errors   []string    `json:"errors"`
}

func Success() *SimpleResponse {
	return &SimpleResponse{
		Success: true,
	}
}

func SuccessData(data interface{}) *SimpleResponse {
	return &SimpleResponse{
		Success: true,
		Data:    data,
	}
}

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
