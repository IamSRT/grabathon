package response

type ApiResponse struct {
	ApiMessage string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
