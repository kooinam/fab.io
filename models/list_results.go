package models

// ListResults used to wrap multiple record results from query
type ListResults struct {
	list   *List
	err    error
	status string
}

// MakeListResults used to instantiate list results
func MakeListResults() *ListResults {
	results := &ListResults{}

	return results
}

// StatusError used to check if operation is errored
func (results *ListResults) StatusError() bool {
	return results.status == StatusError
}

// StatusSuccess used to check if operation is success
func (results *ListResults) StatusSuccess() bool {
	return results.status == StatusSuccess
}

// List used to retrieve result's list
func (results *ListResults) List() *List {
	return results.list
}

// Set used to set result's list and err
func (results *ListResults) Set(list *List, err error) {
	results.list = list
	results.err = err

	if err != nil {
		results.status = StatusError
	} else {
		results.status = StatusSuccess
	}
}

func (results *ListResults) Error() error {
	return results.err
}
