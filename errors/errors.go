package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
	ErrorTypeNotFound       = ErrorType{"not-found"}
)

type Code int

const (
	UnknownCode Code = -1
)

type Error struct {
	code      Code
	errorType ErrorType
	err       error
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Code() Code {
	return e.code
}

func (e Error) ErrorType() ErrorType {
	return e.errorType
}

func NewUnknownError(code Code, err error) Error {
	return Error{
		code:      code,
		errorType: ErrorTypeUnknown,
		err:       err,
	}
}

func NewInvalidFileError(code Code, err error) Error {
	return Error{
		code:      code,
		errorType: ErrorTypeAuthorization,
		err:       err,
	}
}

func NewIncorrectInputError(code Code, err error) Error {
	return Error{
		code:      code,
		errorType: ErrorTypeIncorrectInput,
		err:       err,
	}
}

func NewNotFoundError(code Code, err error) Error {
	return Error{
		code:      code,
		errorType: ErrorTypeNotFound,
		err:       err,
	}
}