package constants

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Result  interface{} `json:"result"`
}

func (r *Response) Error() string {
	return r.Message
}

// success
func NewSuccess(message string, data any) error {
	return &Response{
		Message: message,
		Code:    200,
		Result:  data,
	}
}

// created
func NewCreated(message string, data any) error {
	return &Response{
		Message: message,
		Code:    201,
		Result:  data,
	}
}

// not found error
func NewNotFoundError(message string) error {
	return &Response{
		Message: message,
		Code:    404,
	}
}

// bad request error
func NewBadRequestError(message string) error {
	return &Response{
		Message: message,
		Code:    400,
	}
}

// unauthorized error
func NewUnauthorizedError(message string) error {
	return &Response{
		Message: message,
		Code:    401,
	}
}

// forbidden error
func NewForbiddenError(message string) error {
	return &Response{
		Message: message,
		Code:    403,
	}
}

// internal server error
func NewInternalServerError(message string) error {
	return &Response{
		Message: message,
		Code:    500,
	}
}
