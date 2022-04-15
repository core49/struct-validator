package struct_validator

import (
	"reflect"
)

// Field represents a field from a struct but can also be used separately
type Field struct {
	Name  string
	Value any
	Tag   string
}

// structWalker walks through a struct and collects all struct tag's
// and returns them as a slice of Field
func structWalker(s interface{}) []Field {
	fields := make([]Field, 0)

	sv := reflect.ValueOf(s)
	st := reflect.TypeOf(s)

	for i := 0; i < st.NumField(); i++ {
		rv := sv.Field(i)
		rt := st.Field(i)

		if !rt.IsExported() {
			continue
		}

		switch rv.Kind() {
		case reflect.Struct:
			fields = append(fields, structWalker(rv.Interface())...)
		case reflect.Slice:
			if nt := reflect.ValueOf(rt); nt.Kind() == reflect.Struct {
				sc := reflect.ValueOf(rv.Interface())
				for j := 0; j < sc.Len(); j++ {
					fields = append(fields, structWalker(sc.Index(j).Interface())...)
				}
			}
		default:
			if rt.Tag.Get(validationTag) != "" {
				fields = append(fields, Field{Name: rt.Name, Value: rv.Interface(), Tag: rt.Tag.Get(validationTag)})
			}
		}
	}

	return fields
}
