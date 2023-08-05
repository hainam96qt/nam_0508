package error

type XError struct {
	ErrorMessage string `json:"error"`
}

func (e *XError) Error() string {
	return e.ErrorMessage
}

func NewXError(s string) XError {
	return XError{
		ErrorMessage: s,
	}
}
