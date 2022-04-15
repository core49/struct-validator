package main

import (
	"fmt"
	structValidator "github.com/core49/struct-validator"
)

// FormInput is an example struct for this example
type FormInput struct {
	FirstName string `validate:"alpha"`
	LastName  string `validate:"alpha"`
	Age       uint8  `validate:"numeric"`
}

// main runs the simple-struct example and will print the following result
// valid: false
// one or more field is invalid [LastName 1 ^[a-zA-Z]+$]
// field: LastName value: 1 rule: ^[a-zA-Z]+$
func main() {

	// Create a new Validator
	validator := structValidator.New()

	formInput := FormInput{
		FirstName: "John",
		LastName:  "1",
		Age:       42,
	}

	// check if the struct is valid
	valid, err := validator.Struct(formInput)

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
