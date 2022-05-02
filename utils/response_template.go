package utils

type Response struct {
	Success 	bool 			`json:"success"`
	Message 	string 			`json:"message"`
	Data 		interface{} 	`json:"data,omitempty"`
	Meta 		Meta			`json:"meta,omitempty"`
}

type Meta struct {
	Total int64 `json:"total"`
	Limit int64 `json:"limit"`
	Skip int64 `json:"skip"`
}

func NewSuccessResponse(data interface{}, meta Meta) *Response {
	return &Response{
		Success: true,
		Message: "Success",
		Data: data,
		Meta: meta,
	}
}

func NewFailureResponse(message string, meta Meta) *Response {
	return &Response{
		Success: false,
		Message: message,
		Meta: meta,
	}
}