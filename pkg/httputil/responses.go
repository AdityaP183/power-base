package httputil

// Response represents a standard API response
type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Meta    any  `json:"meta,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// NewResponse creates a new successful response
func NewResponse(data any, meta any) Response {
	return Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(errMsg string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Error:   errMsg,
	}
}
