package storage

import "errors"

var (
	ErrKaboom             = errors.New("Implement me! Kaboom!")
	ErrRecordIdOutOfBands = errors.New("Database Id can't be 0 for certs")
)
