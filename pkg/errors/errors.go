package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common error types
var (
	ErrNotFound       = NewAppError(http.StatusNotFound, "Resource not found", nil)
	ErrInvalidInput   = NewAppError(http.StatusBadRequest, "Invalid input", nil)
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", nil)
	ErrForbidden      = NewAppError(http.StatusForbidden, "Forbidden", nil)
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error", nil)
	ErrDatabase       = NewAppError(http.StatusInternalServerError, "Database error", nil)
	ErrElasticsearch  = NewAppError(http.StatusInternalServerError, "Elasticsearch error", nil)
	ErrTelegramAPI    = NewAppError(http.StatusInternalServerError, "Telegram API error", nil)
)

// WrapError wraps an existing error with additional context
func WrapError(err error, message string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}
