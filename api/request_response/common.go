package request_response

type ApiResponse struct {
	ApiMessage string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
