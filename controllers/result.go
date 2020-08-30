package controllers

// StatusError used to represent errored request
const StatusError = "error"

// StatusUnauthorized used to represent unauthorized request
const StatusUnauthorized = "unauthorized"

// StatusSuccess used to represent success request
const StatusSuccess = "success"

// StatusInternalServerError used to represent unhandled exception request
const StatusInternalServerError = "internalservererror"

// Result used to wrap request result
type Result struct {
	status  string
	err     error
	content interface{}
}

// makeResult used to instantiate result
func makeResult() *Result {
	result := &Result{}

	return result
}

// Set used to set content, status and error
func (result *Result) Set(content interface{}, status string, err error) {
	result.content = content
	result.status = status
	result.err = err
}

// StatusError used to check if request is errored
func (result *Result) StatusError() bool {
	return result.status == StatusError
}

// StatusSuccess used to check if request is success
func (result *Result) StatusSuccess() bool {
	return result.status == StatusSuccess
}

// StatusUnauthorized used to check if request is unauthorized
func (result *Result) StatusUnauthorized() bool {
	return result.status == StatusUnauthorized
}

// Content used to retrieve result's content
func (result *Result) Content() interface{} {
	return result.content
}

// Status used to retrieve result's status
func (result *Result) Status() string {
	return result.status
}

// ErrorMessage used to retrieve result's error message
func (result *Result) ErrorMessage() string {
	if result.err == nil {
		return ""
	}

	return result.err.Error()
}
