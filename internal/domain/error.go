package domain

import "errors"

var (
	ErrorEmptyBody     = errors.New("empty body!")
	ErrorEmptyPassword = errors.New("empty password!")
	ErrorNotFound      = errors.New("entity was not found!")
	ErrorWrongPassword = errors.New("wrong password!")
)
