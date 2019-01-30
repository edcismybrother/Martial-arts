package code

import "errors"

// ErrTYPENOTEXSITED err
var (
	ErrTYPENOTEXSITED = errors.New("the type of this msg does not exist")
	ErrWRONGFORMAT    = errors.New("Wrong message format")
)
