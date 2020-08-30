package models

// ListResults used to wrap multiple record results from query
type ListResults struct {
	items  []Modellable
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

// Items used to retrieve result's all items
func (results *ListResults) Items() []Modellable {
	return results.items
}

// Set used to set result's items and err
func (results *ListResults) Set(items []Modellable, err error) {
	results.items = items
	results.err = err

	if err != nil {
		results.status = StatusError
	} else {
		results.status = StatusSuccess
	}
}
