package struct_validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrValidationResultError(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		fields := []ErrInvalidFields{{
			fieldName: "Answer",
			value:     "forty-two",
			rule:      "^[0-9]*$",
		}}

		err := ErrValidationResult{
			msg:    ErrValidationResultString,
			fields: fields,
		}

		assert.EqualError(t, err, "one or more field is invalid [Answer forty-two ^[0-9]*$]")
		assert.Equal(t, fields, err.Fields())
	})
}

func TestErrInvalidFieldsError(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		err := ErrInvalidFields{
			fieldName: "Answer",
			value:     42,
			rule:      "^[0-9]$",
		}

		assert.EqualError(t, err, "Answer 42 ^[0-9]$")
		assert.Equal(t, err.fieldName, err.Name())
		assert.Equal(t, err.value, err.Value())
		assert.Equal(t, err.rule, err.Rule())
	})
}

func TestErrRuleAlreadyExistsError(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		err := ErrRuleAlreadyExists{
			msg: ErrRuleAlreadyExistsString,
			rule: rule{
				name:       "Name",
				expression: "^[a-zA-Z]+$",
			},
		}

		assert.EqualError(t, err, "rule name already exists {name:Name expression:^[a-zA-Z]+$ rules:<nil>}")
	})
}

func TestErrRuleDoesNotCompile(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		err := ErrRuleDoesNotCompile{
			msg: ErrRuleDoesNotCompileString,
			rule: rule{
				name:       "Test42",
				expression: "^\\d$",
				rules:      nil,
			},
		}

		assert.EqualError(t, err, "unable to compile rule {name:Test42 expression:^\\d$ rules:<nil>}")
	})
}

func TestErrPanicRuleNotDefinedError(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		err := ErrPanicRuleNotDefined{
			msg:       ErrPanicRuleNotDefinedString,
			fieldName: "name",
			ruleName:  "alpha",
		}

		assert.EqualError(t, err, "the called rule does not exist name alpha")
	})
}
