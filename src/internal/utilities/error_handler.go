package utilities

import (
	"errors"
	"net/http"
)

// ErrorHandler provides a fluent interface for handling errors in HTTP responses
type ErrorHandler struct {
	w        http.ResponseWriter
	err      error
	messages []string
	stack    string
	code     int
}

// Error creates a new ErrorHandler instance from a string error message
func Error(msg string, w http.ResponseWriter) *ErrorHandler {
	return &ErrorHandler{
		w:        w,
		err:      errors.New(msg),
		messages: []string{msg},
		stack:    "default-stack",
		code:     http.StatusInternalServerError,
	}
}

// FromError creates a new ErrorHandler from an existing error
func FromError(err error, w http.ResponseWriter) *ErrorHandler {
	if err == nil {
		err = errors.New("unknown error")
	}
	return &ErrorHandler{
		w:        w,
		err:      err,
		messages: []string{err.Error()},
		stack:    "default-stack",
		code:     http.StatusInternalServerError,
	}
}

// NewErrorHandler creates a new ErrorHandler instance
func NewErrorHandler(w http.ResponseWriter, err error) *ErrorHandler {
	return &ErrorHandler{
		w:        w,
		err:      err,
		messages: []string{err.Error()},
		stack:    "default-stack",
		code:     http.StatusInternalServerError,
	}
}

// Msg adds custom messages to the error, appending the original error message
func (e *ErrorHandler) Msg(messages ...string) *ErrorHandler {
	if len(messages) > 0 {
		modifiedMessages := make([]string, len(messages))
		for i, msg := range messages {
			// Append the original error message to each custom message
			modifiedMessages[i] = msg + ": " + e.err.Error()
		}
		e.messages = modifiedMessages
	}
	return e
}

// Stack sets the module/stack name for the error
func (e *ErrorHandler) Stack(name string) *ErrorHandler {
	e.stack = name
	return e
}

// Code sets the HTTP status code for the error response and sends the response
func (e *ErrorHandler) Code(statusCode int) {
	e.code = statusCode
	e.Send()
}

// Send finalizes and sends the error response
func (e *ErrorHandler) Send() {
	if e.w != nil {
		SendError(e.w, e.messages, e.stack, e.code)
	}
}

// --- HTTP Status Code Methods ---

// BadRequest sets status code to 400 and sends the response
func (e *ErrorHandler) BadRequest() {
	e.code = http.StatusBadRequest
	e.Send()
}

// Unauthorized sets status code to 401 and sends the response
func (e *ErrorHandler) Unauthorized() {
	e.code = http.StatusUnauthorized
	e.Send()
}

// Forbidden sets status code to 403 and sends the response
func (e *ErrorHandler) Forbidden() {
	e.code = http.StatusForbidden
	e.Send()
}

// NotFound sets status code to 404 and sends the response
func (e *ErrorHandler) NotFound() {
	e.code = http.StatusNotFound
	e.Send()
}

// MethodNotAllowed sets status code to 405 and sends the response
func (e *ErrorHandler) MethodNotAllowed() {
	e.code = http.StatusMethodNotAllowed
	e.Send()
}

// Conflict sets status code to 409 and sends the response
func (e *ErrorHandler) Conflict() {
	e.code = http.StatusConflict
	e.Send()
}

// TooManyRequests sets status code to 429 and sends the response
func (e *ErrorHandler) TooManyRequests() {
	e.code = http.StatusTooManyRequests
	e.Send()
}

// InternalServerError sets status code to 500 and sends the response
func (e *ErrorHandler) InternalServerError() {
	e.code = http.StatusInternalServerError
	e.Send()
}
