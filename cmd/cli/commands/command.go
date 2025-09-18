package commands

import "errors"

var (
	ErrInvalidArgument error = errors.New("given argument is not valid")
)

type Command interface {
	Execute(args ...string) error
}
