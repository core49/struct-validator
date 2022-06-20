package struct_validator

import (
	"fmt"
)

// validationTag contains the struct tag which is used for this package
const validationTag = "validate"

// Validator is an instance struct of this package
// rules are the loaded rules in the Validator instance
// config is a pointer to the config defined for this Validator instance
type Validator struct {
	rules  rules
	config *config
}

// New returns a new Validator instance and accepts multiple ConfigOption's to configure the instance
// and returns a pointer to a Validator instance
func New(opts ...ConfigOption) (v *Validator) {
	c := &config{}
	for _, optionFunc := range opts {
		optionFunc(c)
	}

	v = &Validator{config: c}
	v.rules = make(rules)

	v.loadBuiltIn()

	return
}

// RegisterRule can be used to define custom rule's
// and will return an error (ErrRuleAlreadyExists|ErrRuleDoesNotCompile|nil)
func (v *Validator) RegisterRule(name, expr string) (err error) {
	r := &rule{
		name:       name,
		expression: expr,
	}

	if er, ok := v.rules[r.name]; ok {
		return ErrRuleAlreadyExists{
			msg:  ErrRuleAlreadyExistsString,
			rule: *er,
		}
	}

	err = r.validate()
	if err != nil {
		return ErrRuleDoesNotCompile{
			msg:  ErrRuleDoesNotCompileString,
			rule: *r,
		}
	}

	v.rules[r.name] = r
	return
}

// Struct is used to validate a struct
// and will return a bool and an error (ErrInvalidFields|nil)
func (v *Validator) Struct(s interface{}) (valid bool, err error) {
	fields := structWalker(s)
	return v.validate(fields...)
}

// Field is used to validate a Field
// and will return a bool and an error (ErrInvalidFields|nil)
func (v *Validator) Field(field Field) (valid bool, err error) {
	return v.validate(field)
}

// loadBuiltIn is used to load the built-in ruleset
func (v *Validator) loadBuiltIn() {
	if v.config.DisableBuiltIn {
		return
	}

	for i := range builtInRules {
		regex := builtInRules[i]

		err := regex.validate()
		if err != nil {
			panic(err)
		}

		v.rules[regex.name] = &regex
	}
}

// validate is used to run over all the Field's and check if they are valid
// and will return a bool and an error (ErrInvalidFields|nil)
func (v *Validator) validate(fields ...Field) (valid bool, err error) {
	var invalidField []ErrInvalidFields

	for _, field := range fields {
		if _, ok := v.rules[field.Tag]; !ok {
			if v.config.PanicOnNotDefinedRules {
				panic(ErrPanicRuleNotDefined{
					msg:       ErrPanicRuleNotDefinedString,
					fieldName: field.Name,
					ruleName:  field.Tag,
				})
			}
			continue
		}

		regex := v.rules[field.Tag].rules
		value := fmt.Sprintf("%v", field.Value)
		fieldValid := regex.MatchString(value)

		if !fieldValid {
			invalidField = append(invalidField, ErrInvalidFields{
				fieldName: field.Name,
				value:     field.Value,
				rule:      regex.String(),
			})
		}
	}

	if len(invalidField) > 0 {
		return false, ErrValidationResult{
			msg:    ErrValidationResultString,
			fields: invalidField,
		}
	}

	return true, nil
}
