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

func Fail() *SimpleResponse {
	return &SimpleResponse{
		Success: false,
	}
}

func Errored(errors ...error) *SimpleResponse {
	e := make([]string, len(errors))
	for _, err := range errors {
		e = append(e, err.Error())
	}
	return &SimpleResponse{
		Success: false,
		Errors:  e,
	}
}
