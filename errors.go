package struct_validator

import "fmt"

var (
	ErrValidationResultString    = "one or more field is invalid"
	ErrRuleAlreadyExistsString   = "rule name already exists"
	ErrRuleDoesNotCompileString  = "unable to compile rule"
	ErrPanicRuleNotDefinedString = "the called rule does not exist"
)

// ErrValidationResult is an error struct which contains a slice of ErrInvalidFields
// msg contains the ErrValidationResultString
// fields contains a slice of ErrInvalidFields
type ErrValidationResult struct {
	msg    string
	fields []ErrInvalidFields
}

// Error returns the msg and fields as string of ErrValidationResult
func (e ErrValidationResult) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.fields)
}

// Fields returns a slice of ErrInvalidFields
func (e ErrValidationResult) Fields() []ErrInvalidFields {
	return e.fields
}

// ErrInvalidFields is an error struct which contains more information about an invalid field
// fieldName contains the name of the invalid field
// value contains the value of the invalid field
// ruleName contains the regex of the used ruleName for this invalid field
type ErrInvalidFields struct {
	fieldName string
	value     interface{}
	rule      string
}

// Error returns the ErrInvalidFields as a string
func (e ErrInvalidFields) Error() string {
	return fmt.Sprintf("%s %v %s", e.fieldName, e.value, e.rule)
}

// Name returns the fieldName of the struct ErrInvalidFields
func (e ErrInvalidFields) Name() string {
	return e.fieldName
}

// Value returns the value of the struct ErrInvalidFields
func (e ErrInvalidFields) Value() interface{} {
	return e.value
}

// Rule returns the used ruleName for this struct ErrInvalidFields
func (e ErrInvalidFields) Rule() string {
	return e.rule
}

// ErrRuleAlreadyExists is an error struct which returns the already defined rule
// msg contains the ErrRuleAlreadyExistsString
// rule contains the rule which already exists
type ErrRuleAlreadyExists struct {
	msg  string
	rule rule
}

// Error returns the msg string of ErrRuleAlreadyExists
func (e ErrRuleAlreadyExists) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.rule)
}

// ErrRuleDoesNotCompile is an error struct which contains the ruleName which could not be compiled
// msg contains the ErrRuleDoesNotCompile
// rule contains the rule which does not compile
type ErrRuleDoesNotCompile struct {
	msg  string
	rule rule
}

// Error returns the msg string of ErrRuleDoesNotCompile
func (e ErrRuleDoesNotCompile) Error() string {
	return fmt.Sprintf("%s %+v", e.msg, e.rule)
}

// ErrPanicRuleNotDefined is an error struct which returns the not defined filed rule
// msg contains the ErrPanicRuleNotDefinedString
// fieldName contains the name of the field with a not defined rule
// ruleName contains the name of the not defined rule
type ErrPanicRuleNotDefined struct {
	msg       string
	fieldName string
	ruleName  string
}

// Error returns the msg string of ErrPanicRuleNotDefined
func (e ErrPanicRuleNotDefined) Error() string {
	return fmt.Sprintf("%s %s %s", e.msg, e.fieldName, e.ruleName)
}
