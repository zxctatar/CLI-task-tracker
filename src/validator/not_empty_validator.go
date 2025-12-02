package validator

import "errors"

type NotEmptyValidator struct {}

func (nev *NotEmptyValidator) Check(args []string) error {
	if len(args) == 0 {
		return errors.New("empty input")
	}
	return nil
}