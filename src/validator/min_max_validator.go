package validator

import (
	"errors"
)

type MinMaxValidator struct {
	min int
	max int
}

func NewMinMaxValidator(min int, max int) *MinMaxValidator {
	if min > max {
		temp := min
		min = max
		max = temp
	}
	return &MinMaxValidator{min, max}
}

func (mmv *MinMaxValidator) Check(args []string) error {
	if len(args) > mmv.max || len(args) < mmv.min {
		return errors.New("invalid number of arguments")
	}

	return nil
}
