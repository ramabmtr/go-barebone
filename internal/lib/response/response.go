package response

type Response struct {
	Message *string `json:"message,omitempty"`
	Data    any     `json:"data,omitempty"`
}

func Message(message string) *Response {
	return &Response{
		Message: &message,
	}
}

func Success(data any, envelope ...string) *Response {
	if len(envelope) > 0 {
		data = map[string]any{
			envelope[0]: data,
		}
	}

	return &Response{
		Data: data,
	}
}

func Error(err error) *Response {
	msg := err.Error()
	return &Response{
		Message: &msg,
	}
}
