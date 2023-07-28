package client

import "github.com/pkg/errors"

var (
	ErrNotFound = errors.New("failed to get resource")
)
