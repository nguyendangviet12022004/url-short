package exception

import "net/http"

type CustomError struct {
	Code      int    `json:"-"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func NewCustomError(code int, errorCode string, message string) *CustomError {
	return &CustomError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   message,
	}
}

var (
	ErrUserNotFound         = NewCustomError(http.StatusNotFound, "USER_NOT_FOUND", "user not found")
	ErrUserExists           = NewCustomError(http.StatusBadRequest, "USER_EXISTS", "user already exists")
	ErrInvalidRequestBody   = NewCustomError(http.StatusBadRequest, "INVALID_REQUEST_BODY", "invalid request body")
	ErrInvalidRequestParams = NewCustomError(http.StatusBadRequest, "INVALID_REQUEST_PARAMS", "invalid request params")
	ErrInternalServer       = NewCustomError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error")
	ErrSql                  = NewCustomError(http.StatusInternalServerError, "SQL_ERROR", "sql error")
	ErrBadCredential        = NewCustomError(http.StatusBadRequest, "BAD_CREDENTIAL", "bad credential")
)
