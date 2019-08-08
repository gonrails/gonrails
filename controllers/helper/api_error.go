package helper

type APIError struct {
	Code         int    `json:"-"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Request      string `json:"request"`
}

//type error interface {
//	Error() string
//}

func (e *APIError) Error() string {
	return e.ErrorMessage
}

func NewAPIError(code, errorCode int, errorMessage, request string) *APIError {
	return &APIError{
		Code:         code,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Request:      request,
	}
}
