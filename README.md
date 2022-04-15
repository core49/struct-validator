# struct-validator
[![codecov](https://codecov.io/gh/core49/struct-validator/branch/main/graph/badge.svg?token=X1WRWVFOZG)](https://codecov.io/gh/core49/struct-validator)

A simple validator to check defined fields in a struct with regular expression.

Contrary to the [go-playground/validator](https://github.com/go-playground/validator) package, which tries to give you as many options as possible, struct-validator only requires a regex to fulfill its tasks.

# Installation

Use **go get**.

        go get github.com/core49/struct-validator

Import struct-validator into your code.

        import "github.com/core49/struct-validator"

# Return values



The package returns two values. The first one is of type **boolean** and the second one is of type [**error** (ErrValidationResult)](https://pkg.go.dev/github.com/core49/struct-validator#ErrValidationResult).

```go
	// check if the struct is valid
	valid, err := validator.Struct(formInput)
```

Furthermore the error value contains extended information about the executed validation. Struct-validator is able to return every checked field and the corresponding rule it was evaluated with.

```go
	if err != nil {
		// detailed information to process
		result := err.(structValidator.ErrValidationResult)
		for _, field := range result.Fields() {
			fmt.Printf("field: %s value: %v rule: %s", field.Name(), field.Value(), field.Rule())
		}
	}
```
