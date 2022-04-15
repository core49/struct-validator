package struct_validator

import "regexp"

// rules contains multiple rule's
// key: name of rule
// value: &rule
type rules map[string]*rule

// rule contains all needed information for a new rule
// name contains the name of the rule
// expression contains a regular expression string
// rules has a pointer to Regexp, which is the representation of a compiled regular expression.
type rule struct {
	name       string
	expression string
	rules      *regexp.Regexp
}

// validate is used to check if the regular expression of a rule is valid / compiles
// and returns an error
func (r *rule) validate() (err error) {
	r.rules, err = regexp.Compile(r.expression)
	return
}
