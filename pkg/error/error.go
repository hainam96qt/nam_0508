package error

type XError struct {
	ErrorMessage string `json:"error"`
	HTTPStatus   int    `json:"-"`
}

func (e XError) Error() string {
	return e.ErrorMessage
}

func (e XError) GetHTTPStatus() int {
	return e.HTTPStatus
}

func NewXError(s string, httpStatus int) XError {
	return XError{
		ErrorMessage: s,
		HTTPStatus:   httpStatus,
	}
}
