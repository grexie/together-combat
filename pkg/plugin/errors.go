package plugin

import (
	"fmt"

	"github.com/graphql-go/graphql/gqlerrors"
)

type Error struct {
	source error
	extensions map[string]any
}

var _ gqlerrors.ExtendedError = &Error{}

func (e *Error) Error() string {
	return e.source.Error()
}

func (e *Error) Source() error {
	return e.source
}

func (e *Error) Extensions() map[string]any {
	return e.extensions
}

func NewError(httpStatus int, code string, message string, args ...any) error {
	source := fmt.Errorf(message, args...)
	err := &Error{
		source: source,
		extensions: map[string]any{
			"httpStatus": httpStatus,
			"code": code,
			"message": source.Error(),
		},
	}
	return err
}

func NewInvalidTokenError(message string, args ...any) error {
	return NewError(401, "invalid-token", message, args...)
}

func NewAuthenticationError(message string, args ...any) error {
	return NewError(401, "unauthenticated", message, args...)
}

func NewForbiddenError(message string, args ...any) error {
	return NewError(403, "forbidden", message, args...)
}

func NewBadRequestError(message string, args ...any) error {
	return NewError(400, "bad-request", message, args...)
}

func NewNotFoundError(message string, args ...any) error {
	return NewError(404, "not-found", message, args...)
}

func NewInvalidError(message string, args ...any) error {
	return NewError(422, "invalid", message, args...)
}

func NewServerError(source error) error {
	return NewError(500, "server-error", "An unknown application error occurred and the server was unable to process your request.")
}

func NewNotImplementedError(source error) error {
	return NewError(500, "server-error", "Not implemented.")
}