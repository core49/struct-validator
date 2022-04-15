package main

import (
	structValidator "github.com/core49/struct-validator"
)

// main runs the field example with defined options and will print the following result
// panic: the called rule does not exist FirstName alpha
func main() {

	// Create a new Validator
	// Set the following options
	// -> disable the built-in rules
	// -> panic if called rule is not defined
	validator := structValidator.New(structValidator.OptionDisableBuiltInRules(true), structValidator.OptionPanicOnNotDefinedRules(true))

	f := structValidator.Field{
		Name:  "FirstName",
		Value: "John",
		Tag:   "alpha",
	}

	// will result in the configured panic since the built-in options are disabled
	_, _ = validator.Field(f)
}
