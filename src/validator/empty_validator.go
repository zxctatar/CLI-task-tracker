package validator

import "errors"

type EmptyValudator struct{}

func (ev *EmptyValudator) Check(args []string) error {
	if len(args) > 0 {
		return errors.New("the command works without arguments")
	}
	return nil
}
