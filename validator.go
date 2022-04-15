package struct_validator

import (
	"fmt"
)

const validationTag = "validate"

type Validator struct {
	rules  rules
	config *config
}

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

func (v *Validator) Struct(s interface{}) (valid bool, err error) {
	fields := structWalker(s)
	return v.validate(fields...)
}

func (v *Validator) Field(field Field) (valid bool, err error) {
	return v.validate(field)
}

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

func (v *Validator) validate(fields ...Field) (valid bool, err error) {
	var invalidField []ErrInvalidFields

	for _, field := range fields {
		if _, ok := v.rules[field.Tag]; !ok {
			if v.config.PanicOnNotDefinedRules {
				panic(ErrPanicRuleNotDefined{
					msg:   ErrPanicRuleNotDefinedString,
					field: field.Name,
					rule:  field.Tag,
				})
			}
			continue
		}

		regex := v.rules[field.Tag].rules
		value := fmt.Sprintf("%v", field.Value)
		fieldValid := regex.MatchString(value)

		if !fieldValid {
			invalidField = append(invalidField, ErrInvalidFields{
				field: field.Name,
				value: field.Value,
				rule:  regex.String(),
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
