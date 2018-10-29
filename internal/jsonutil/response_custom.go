package jsonutil

// Success return a success json response
// equavalent to {"success" : true}
func Success() *DefaultResponse {
	return &DefaultResponse{
		Success: true,
	}
}

// SuccessData return a success json response with
// extra data informations
// example: {"success" : true, "data": {"name": "r1",
// "test": true}}
func SuccessData(data interface{}) *DefaultResponse {
	return &DefaultResponse{
		Success: true,
		Data:    data,
	}
}

// Fail return a non success reponse with success = false
// adding all passed errors to the response
// example {"success" : false, "errors": ["can't find my self",
// "where am i "]}
func Fail(errors ...error) *DefaultResponse {
	errs := make([]string, 0, len(errors))
	for _, err := range errors {
		errs = append(errs, err.Error())
	}
	return &DefaultResponse{
		Success: false,
		Errors:  errs,
	}
}

// FailData .
func FailData(data interface{}, errors ...error) *DefaultResponse {
	errs := make([]string, 0, len(errors))
	for _, err := range errors {
		errs = append(errs, err.Error())
	}
	return &DefaultResponse{
		Success: false,
		Data:    data,
		Errors:  errs,
	}
}
