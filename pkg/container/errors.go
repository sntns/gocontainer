package container

import "errors"

var (
	ErrInvalidCopyName      = errors.New("invalid copy name (format: <file>[:<destination>])")
	ErrorInvalidLabelFormat = errors.New("invalid label (format: <key>=<value>)")
)
