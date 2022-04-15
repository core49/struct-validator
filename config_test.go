package struct_validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionDisableBuiltInRules(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		c := &config{}
		co := OptionDisableBuiltInRules(true)
		co(c)

		assert.True(t, c.DisableBuiltIn)

		co = OptionDisableBuiltInRules(false)
		co(c)

		assert.False(t, c.DisableBuiltIn)
	})
}

func TestOptionPanicOnNotDefinedRules(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		c := &config{}
		co := OptionPanicOnNotDefinedRules(true)
		co(c)

		assert.True(t, c.PanicOnNotDefinedRules)

		co = OptionPanicOnNotDefinedRules(false)
		co(c)

		assert.False(t, c.PanicOnNotDefinedRules)
	})
}
