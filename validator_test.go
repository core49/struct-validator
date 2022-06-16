package struct_validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("WithoutOptions", func(t *testing.T) {
			v := New()

			assert.NotNil(t, v.config)
			assert.NotNil(t, v.rules)
			assert.Equal(t, len(builtInRules), len(v.rules))
		})
		t.Run("OptionDisableBuiltInRules", func(t *testing.T) {
			v := New(OptionDisableBuiltInRules(true))

			assert.True(t, v.config.DisableBuiltIn)
		})
		t.Run("OptionPanicOnNotDefinedRules", func(t *testing.T) {
			v := New(OptionPanicOnNotDefinedRules(true))

			assert.True(t, v.config.PanicOnNotDefinedRules)
		})
	})
}

func TestValidatorRegisterRule(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		v := New(OptionDisableBuiltInRules(true))

		err := v.RegisterRule("test", "^test$")

		assert.NoError(t, err)
		assert.Equal(t, 1, len(v.rules))
	})
	t.Run("FAIL", func(t *testing.T) {
		t.Run("ErrRuleAlreadyExists", func(t *testing.T) {
			v := New(OptionDisableBuiltInRules(true))

			err := v.RegisterRule("test", "^test$")

			assert.NoError(t, err)
			assert.Equal(t, 1, len(v.rules))

			err = v.RegisterRule("test", "^test$")

			assert.EqualError(t, err, ErrRuleAlreadyExistsString+fmt.Sprintf(" %+v", *v.rules["test"]))
			assert.Equal(t, 1, len(v.rules))
		})
		t.Run("ErrRuleDoesNotCompile", func(t *testing.T) {
			v := New(OptionDisableBuiltInRules(true))

			r := rule{name: "test", expression: "["}
			err := v.RegisterRule(r.name, r.expression)

			assert.EqualError(t, err, ErrRuleDoesNotCompileString+fmt.Sprintf(" %+v", r))
			assert.Equal(t, 0, len(v.rules))
		})
	})
}

func TestValidatorStruct(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		type FormInput struct {
			FirstName string `validate:"alpha"`
			LastName  string `validate:"alpha"`
			Age       int    `validate:"numeric"`
		}

		fi := FormInput{
			FirstName: "John",
			LastName:  "Doe",
			Age:       42,
		}

		v := New()
		valid, err := v.Struct(fi)

		assert.NoError(t, err)
		assert.True(t, valid)
	})
	t.Run("FAIL", func(t *testing.T) {
		type FormInput struct {
			FirstName string `validate:"numeric"`
			LastName  string `validate:"numeric"`
			Age       int    `validate:"numeric"`
		}

		fi := FormInput{
			FirstName: "John",
			LastName:  "Doe",
			Age:       42,
		}

		v := New()
		valid, err := v.Struct(fi)

		assert.EqualError(t, err, "one or more field is invalid [FirstName John ^[-+]?[0-9]+(?:\\.[0-9]+)?$ LastName Doe ^[-+]?[0-9]+(?:\\.[0-9]+)?$]")
		assert.False(t, valid)
	})
}

func TestValidatorField(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		f := Field{
			Name:  "FirstName",
			Value: "John",
			Tag:   "alpha",
		}

		v := New()
		valid, err := v.Field(f)
		assert.NoError(t, err)
		assert.True(t, valid)
	})
	t.Run("FAIL", func(t *testing.T) {
		f := Field{
			Name:  "FirstName",
			Value: "John",
			Tag:   "numeric",
		}

		v := New()
		valid, err := v.Field(f)
		assert.EqualError(t, err, "one or more field is invalid [FirstName John ^[-+]?[0-9]+(?:\\.[0-9]+)?$]")
		assert.False(t, valid)
	})
}

func TestValidatorLoadBuiltIn(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		v := New(OptionDisableBuiltInRules(true))

		assert.Equal(t, 0, len(v.rules))

		v.config.DisableBuiltIn = false
		v.loadBuiltIn()

		assert.Equal(t, len(builtInRules), len(v.rules))
	})
	t.Run("FAIL", func(t *testing.T) {
		v := New(OptionDisableBuiltInRules(true))

		assert.Equal(t, 0, len(v.rules))

		v.config.DisableBuiltIn = false
		builtInRules = append(builtInRules, rule{
			name:       "_test_",
			expression: "[",
		})

		assert.Panics(t, func() { v.loadBuiltIn() })
		assert.Equal(t, len(builtInRules)-1, len(v.rules))

		builtInRules = builtInRules[:len(builtInRules)-1]
	})
}

func TestValidatorValidate(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			f := Field{
				Name:  "FirstName",
				Value: "John",
				Tag:   "alpha",
			}
			v := New()
			valid, err := v.validate(f)

			assert.NoError(t, err)
			assert.True(t, valid)
		})
		t.Run("Invalid", func(t *testing.T) {
			f := Field{
				Name:  "FirstName",
				Value: "John",
				Tag:   "test",
			}
			v := New()
			valid, err := v.validate(f)

			assert.NoError(t, err)
			assert.True(t, valid)
		})
	})
	t.Run("FAIL", func(t *testing.T) {
		t.Run("ErrPanicRuleNotDefined", func(t *testing.T) {
			f := Field{
				Name:  "FirstName",
				Value: "John",
				Tag:   "numeric",
			}
			v := New(OptionDisableBuiltInRules(true), OptionPanicOnNotDefinedRules(true))

			assert.Panics(t, func() { _, _ = v.validate(f) })
		})
		t.Run("ErrInvalidFields", func(t *testing.T) {
			f := Field{
				Name:  "FirstName",
				Value: "John",
				Tag:   "numeric",
			}
			v := New()
			valid, err := v.validate(f)

			assert.EqualError(t, err, "one or more field is invalid [FirstName John ^[-+]?[0-9]+(?:\\.[0-9]+)?$]")
			assert.False(t, valid)
		})
	})
}
