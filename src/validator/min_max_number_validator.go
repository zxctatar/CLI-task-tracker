package validator

import (
	"errors"
	"strconv"
)

type MinMaxNumberValidator struct {
	min int
	max int
}

func NewMinMaxNumberValidator(min int, max int) *MinMaxNumberValidator {
	if min > max {
		temp := min
		min = max
		max = temp
	}
	return &MinMaxNumberValidator{min, max}
}

func (mmv *MinMaxNumberValidator) Check(args []string) error {
	if len(args) > mmv.max || len(args) < mmv.min {
		return errors.New("invalid number of arguments")
	}

	for _, a := range args {
		_, err := strconv.Atoi(a)
		if err != nil {
			return errors.New("argument is not a number")
		}
	}

	return nil
}
