package controllers

// NetworkError encapsulates network error with status code and error message
type NetworkError struct {
	Status int
	Error  string
}
