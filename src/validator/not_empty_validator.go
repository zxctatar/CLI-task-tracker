package validator

import "errors"

type NotEmptyValidator struct{}

func (nev *NotEmptyValidator) Check(args []string) error {
	if len(args) == 0 {
		return errors.New("the command is waiting for arguments")
	}
	return nil
}
