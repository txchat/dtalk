package error

func NewCustomError(e *Error, data interface{}) *Error {
	newErr := &Error{
		code: e.Code(),
		msg:  e.msg,
		data: data,
	}
	return newErr
}
