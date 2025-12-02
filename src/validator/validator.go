package validator

type Validator interface {
	Check(args []string) error
}