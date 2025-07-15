package errors

type AppError struct {
	Code       int    `json:"code"`    // Application-specific code (e.g. 701)
	Message    string `json:"message"` // User-facing message
	HTTPStatus int    `json:"-"`       // HTTP response code (e.g. 400)
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string, httpStatus int) *AppError {
	return &AppError{Code: code, Message: message, HTTPStatus: httpStatus}
}
