package storage

import "errors"

var (
	ErrAlreadyExists = errors.New("cloud already exists")
	ErrNotFound      = errors.New("cloud not found")
)
