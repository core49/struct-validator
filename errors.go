package struct_validator

import "fmt"

var (
	ErrValidationResultString    = "one or more field is invalid"
	ErrRuleAlreadyExistsString   = "rule name already exists"
	ErrRuleDoesNotCompileString  = "unable to compile rule"
	ErrPanicRuleNotDefinedString = "the called rule does not exist"
)

type ErrValidationResult struct {
	msg    string
	fields []ErrInvalidFields
}

func (e ErrValidationResult) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.fields)
}

func (e ErrValidationResult) Fields() []ErrInvalidFields {
	return e.fields
}

type ErrInvalidFields struct {
	field string
	value interface{}
	rule  string
}

func (e ErrInvalidFields) Error() string {
	return fmt.Sprintf("%s %v %s", e.field, e.value, e.rule)
}

func (e ErrInvalidFields) Name() string {
	return e.field
}

func (e ErrInvalidFields) Value() interface{} {
	return e.value
}

func (e ErrInvalidFields) Rule() string {
	return e.rule
}

type ErrRuleAlreadyExists struct {
	msg  string
	rule rule
}

func (e ErrRuleAlreadyExists) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.rule)
}

type ErrRuleDoesNotCompile struct {
	msg  string
	rule rule
}

func (e ErrRuleDoesNotCompile) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.rule)
}

type ErrPanicRuleNotDefined struct {
	msg   string
	field string
	rule  string
}

func (e ErrPanicRuleNotDefined) Error() string {
	return fmt.Sprintf("%s %s %s", e.msg, e.field, e.rule)
}
