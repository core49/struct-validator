package struct_validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// invalidRegex is a string slice containing invalid regular expression
var invalidRegex = []string{"[", "*", "[0-9]++", "FOO\\"}

// TestRuleValidate is used to test the following function
// func (r *rule) validate() (err error)
func TestRuleValidate(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		for _, builtInRule := range builtInRules {
			t.Run(builtInRule.name, func(t *testing.T) {
				err := builtInRule.validate()
				assert.NoError(t, err)
			})
		}
	})
	t.Run("FAIL", func(t *testing.T) {
		for _, regex := range invalidRegex {
			r := rule{
				name:       regex,
				expression: regex,
			}
			t.Run(r.name, func(t *testing.T) {
				err := r.validate()
				assert.Error(t, err)
			})
		}
	})
}
