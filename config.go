package struct_validator

type config struct {
	DisableBuiltIn         bool
	PanicOnNotDefinedRules bool
}

type ConfigOption func(config *config)

func OptionDisableBuiltInRules(v bool) ConfigOption {
	return func(c *config) {
		c.DisableBuiltIn = v
	}
}

func OptionPanicOnNotDefinedRules(v bool) ConfigOption {
	return func(c *config) {
		c.PanicOnNotDefinedRules = v
	}
}
