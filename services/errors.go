package services

import "errors"

var (
	ErrNonUniqueEmail      = errors.New("services: account with the specified email already exist")
	ErrNonUniqueSourceName = errors.New("services: user already have source with the specified name")
	ErrEmptyEmail          = errors.New("services: email must not be empty")
	ErrEmptyPassword       = errors.New("services: password must not be empty")
	ErrEmptySourceName     = errors.New("services: source name must not be empty")
	ErrEmptySourceURL      = errors.New("services: source url must not be empty")
)
