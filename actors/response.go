package actors

// Response represents the response for synchronous task delegation request for action
type Response struct {
	status  int // 0 -> success, 1 -> failed
	message string
}

func makeResponse(status int, message string) *Response {
	response := &Response{
		status:  status,
		message: message,
	}

	return response
}
