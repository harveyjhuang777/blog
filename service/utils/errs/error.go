package errs

import "errors"

type code int

const (
	CodeFieldEmpty code = iota
	CodeFieldInvalid
	CodeNotFound
	CodeConflict
	CodeForbidden
	CodeBadRequest
	CodeServerError
)

func (c code) Name() string {
	names := [...]string{
		"FIELD_EMPTY",
		"FIELD_INVALID",
		"NOT_FOUND",
		"CONFLICT",
		"FORBIDDEN",
		"BAD_REQUEST",
		"SERVER_ERROR",
	}
	return names[c]
}

var (
	ErrNoRequestBody      = errors.New("no request body.")
	ErrRecordFailed       = errors.New("record create failed.")
	ErrInvalidRequest     = errors.New("something wrong in your request.")
	ErrRecordExist        = errors.New("record already exists!")
	ErrValueTooLong       = errors.New("field value too long.")
	ErrFieldNotExist      = errors.New("field value not exists.")
	ErrRecordNotFound     = errors.New("record not found.")
	ErrServerError        = errors.New("server internal error.")
	ErrInvalidTimeFormat  = errors.New("time format worng.")
	ErrTimeOverlap        = errors.New("transaction time overlap.")
	ErrNoAccessToken      = errors.New("no access token.")
	ErrInvalidAccessToken = errors.New("access token invalid.")
	ErrNoAccessRight      = errors.New("no access right.")
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token")
)
