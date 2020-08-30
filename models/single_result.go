package models

// StatusFailed used to represent failed query or operation
const StatusFailed = "failed"

// StatusNotFound used to represent not found query
const StatusNotFound = "notfound"

// StatusSuccess used to represent success query or operation
const StatusSuccess = "success"

// SingleResult used to wrap single record result from query or create operation
type SingleResult struct {
	item   Modellable
	err    error
	status string
}

// MakeSingleResult used to instantiate single result
func MakeSingleResult() *SingleResult {
	result := &SingleResult{}

	return result
}

// Failed used to check if operation is failed
func (result *SingleResult) Failed() bool {
	return result.status == StatusFailed
}

// NotFound used to check if record is not found
func (result *SingleResult) NotFound() bool {
	return result.status == StatusNotFound
}

// Success used to check if operation is success
func (result *SingleResult) Success() bool {
	return result.status == StatusSuccess
}

// Status used to retrieve result's status
func (result *SingleResult) Status() string {
	return result.status
}

func (result *SingleResult) Error() error {
	return result.err
}

// Set used to set result item, error and status
func (result *SingleResult) Set(item Modellable, err error, notFound bool) {
	result.item = item
	result.err = err

	if notFound {
		result.status = StatusNotFound
	} else if err != nil {
		result.status = StatusFailed
	} else {
		result.status = StatusSuccess
	}
}

// ErrorMessage used to retrieve result's error message
func (result *SingleResult) ErrorMessage() string {
	if result.err == nil {
		return ""
	}

	return result.err.Error()
}

// Item used to retrieve result's item
func (result *SingleResult) Item() Modellable {
	return result.item
}
