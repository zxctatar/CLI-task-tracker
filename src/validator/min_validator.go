package validator

import "errors"

type MinValidator struct {
	min int
}

func NewMinValidator(min int) *MinValidator {
	return &MinValidator{min}
}

func (mv *MinValidator) Check(args []string) error {
	if len(args) < mv.min {
		return errors.New("invalid number of arguments")
	}
	return nil
}
