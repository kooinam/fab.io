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

// Failed used to check if operation is failed
func (results *ListResults) Failed() bool {
	return results.status == StatusFailed
}

// Success used to check if operation is success
func (results *ListResults) Success() bool {
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
		results.status = StatusFailed
	} else {
		results.status = StatusSuccess
	}
}
