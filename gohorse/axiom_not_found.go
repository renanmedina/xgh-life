package gohorse

import "errors"

type AxiomNotFoundError struct {
	err error
}

func (e AxiomNotFoundError) Error() string {
	return e.err.Error()
}

func NewAxiomNotFoundError(details string) error {
	return AxiomNotFoundError{
		err: errors.New(details),
	}
}
