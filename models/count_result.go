package models

// CountResult used to wrap count records result from query
type CountResult struct {
	count  int64
	err    error
	status string
}

// MakeCountResult used to instantiate count result
func MakeCountResult() *CountResult {
	result := &CountResult{}

	return result
}

// StatusError used to check if operation is errored
func (result *CountResult) StatusError() bool {
	return result.status == StatusError
}

// StatusSuccess used to check if operation is success
func (result *CountResult) StatusSuccess() bool {
	return result.status == StatusSuccess
}

// Count use to retrieve result's count
func (result *CountResult) Count() int64 {
	return result.count
}

// Set used to set result count and error
func (result *CountResult) Set(count int64, err error) {
	result.count = count
	result.err = err

	if err != nil {
		result.status = StatusError
	} else {
		result.status = StatusSuccess
	}
}

// Status used to retrieve result's status
func (result *CountResult) Status() string {
	return result.status
}
