package controllers

// ResponseError defines the structure of the error response
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseMessage defines the structure of the success message response
type ResponseMessage struct {
	Message string `json:"message"`
}
