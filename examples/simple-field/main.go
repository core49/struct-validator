package main

import (
	"fmt"
	structValidator "github.com/core49/struct-validator"
)

// main runs the field example and will print the following result
// valid: false
// one or more field is invalid [FirstName 42 ^[a-zA-Z]+$]
// field: FirstName value: 42 rule: ^[a-zA-Z]+$
func main() {

	// Create a new Validator
	validator := structValidator.New()

	f := structValidator.Field{
		Name:  "FirstName",
		Value: 42,
		Tag:   "alpha",
	}

	// check if the field is valid
	valid, err := validator.Field(f)

	fmt.Printf("valid: %t\n", valid)

	if err != nil {
		// default error message
		fmt.Println(err.Error())

		// detailed information to process
		result := err.(structValidator.ErrValidationResult)
		for _, field := range result.Fields() {
			fmt.Printf("field: %s value: %v rule: %s", field.Name(), field.Value(), field.Rule())
		}
	}
}
