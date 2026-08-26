package model

import "errors"

var (
	ErrInvalidRecord  = errors.New("invalid record")
	ErrInvalidProfile = errors.New("invalid profile")
	ErrInvalidEvent   = errors.New("invalid event")
	ErrInvalidAudit   = errors.New("invalid audit")
	ErrInvalidStatus  = errors.New("invalid status")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
)
