package exception

import "fmt"

type NotFoundException struct {
	format string
	arg    []any
}

func ThrowNotFoundException(format string, arg ...any) *NotFoundException {
	return &NotFoundException{
		format: format,
		arg:    arg,
	}
}

func (r *NotFoundException) Error() string {
	return fmt.Sprintf(r.format, r.arg...)
}
