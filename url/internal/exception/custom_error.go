package exception

type CustomError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
}

func (e *CustomError) Error() string {
	return e.Message
}

var (
	ErrInternalServer = &CustomError{
		Code:       500,
		Message:    "Internal Server Error",
		HttpStatus: 500,
	}
	ErrBadRequest = &CustomError{
		Code:       400,
		Message:    "Bad Request",
		HttpStatus: 400,
	}
	ErrNotFound = &CustomError{
		Code:       404,
		Message:    "Not Found",
		HttpStatus: 404,
	}
	ErrUnauthorized = &CustomError{
		Code:       401,
		Message:    "Unauthorized",
		HttpStatus: 401,
	}
	ErrForbidden = &CustomError{
		Code:       403,
		Message:    "Forbidden",
		HttpStatus: 403,
	}
	ErrConflict = &CustomError{
		Code:       409,
		Message:    "Conflict",
		HttpStatus: 409,
	}
	ErrTooManyRequests = &CustomError{
		Code:       429,
		Message:    "Too Many Requests",
		HttpStatus: 429,
	}
	ErrServiceUnavailable = &CustomError{
		Code:       503,
		Message:    "Service Unavailable",
		HttpStatus: 503,
	}
	ErrShortUrlNotFound = &CustomError{
		Code:       404,
		Message:    "Short URL not found",
		HttpStatus: 404,
	}
)
