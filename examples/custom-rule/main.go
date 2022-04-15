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
	HairColor string `validate:"hairColor"`
}

// main runs the simple-struct example and will print the following result
// valid: true
func main() {

	// Create a new Validator
	validator := structValidator.New()

	formInput := FormInput{
		FirstName: "John",
		LastName:  "Doe",
		Age:       42,
		HairColor: "blonde",
	}

	// Register custom rule
	// only use the four main hair colors
	registerErr := validator.RegisterRule("hairColor", "^blonde|brunette|red|black$")
	if registerErr != nil {

		// this will occur if there is already a rule registered with this name
		if ruleErr, ok := registerErr.(structValidator.ErrRuleAlreadyExists); ok {
			fmt.Println(ruleErr.Error())
		}

		/// this will occur if there provided regular expression could not be compiled (is not valid)
		if compileErr, ok := registerErr.(structValidator.ErrRuleDoesNotCompile); ok {
			fmt.Println(compileErr.Error())
		}
	}

	// check if the struct is valid
	valid, err := validator.Struct(formInput)

	fmt.Printf("valid: %t\n", valid)

	if err != nil {
		// default error message
		fmt.Println(err.Error())
	}
}
