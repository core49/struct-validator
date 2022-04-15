package struct_validator

// config contains the configured config options
// DisableBuiltIn is true when the built-in ruleset is disabled. Default false
// PanicOnNotDefinedRules is set to true if no panic is desired on an undefined rule in a fieldName. Default false
type config struct {
	DisableBuiltIn         bool
	PanicOnNotDefinedRules bool
}

// ConfigOption is a function pseudonym
type ConfigOption func(config *config)

// OptionDisableBuiltInRules is a ConfigOption to disable the built-in ruleset
func OptionDisableBuiltInRules(v bool) ConfigOption {
	return func(c *config) {
		c.DisableBuiltIn = v
	}
}

// OptionPanicOnNotDefinedRules is a ConfigOption to disable a panic on an undefined ruleName in a fieldName
func OptionPanicOnNotDefinedRules(v bool) ConfigOption {
	return func(c *config) {
		c.PanicOnNotDefinedRules = v
	}
}
