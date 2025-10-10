package domain

import "errors"

// Domain-specific errors
var (
	ErrCharacterNameRequired = errors.New("character name is required")
	ErrInvalidDimensions     = errors.New("character dimensions must be positive")
	ErrEmptyPattern          = errors.New("pattern cannot be empty")
	ErrInvalidFrameCount     = errors.New("frame count does not match character height")
	ErrCharacterNotFound     = errors.New("character not found")
	ErrInvalidFrameName      = errors.New("frame name cannot be empty")
	ErrInvalidPattern        = errors.New("invalid pattern")
)
