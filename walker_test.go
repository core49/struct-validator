//nolint:jscpd
package struct_validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// user is a struct used for multiple test
type user struct {
	FirstName string `validate:"alpha"`
	LastName  string `validate:"alpha"`
	Age       uint8  `validate:"number"`
	Email     string `validate:"email"`
}

// TestValidatorSliceWalker is used to test the following function
// func structWalker(s interface{}) []Field
func TestValidatorSliceWalker(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("Struct", func(t *testing.T) {
			s := user{
				"John",
				"Doe",
				42,
				"john.doe@domain.tld",
			}

			structFields := []map[string]string{
				{"FirstName": "alpha"},
				{"LastName": "alpha"},
				{"Age": "number"},
				{"Email": "email"},
			}

			fields := structWalker(s)
			for i, field := range fields {
				for key, value := range structFields[i] {
					assert.Equal(t, key, field.Name)
					assert.Equal(t, value, field.Tag)
				}
			}

			assert.Equal(t, len(structFields), len(fields))
		})
		t.Run("MultidimensionalStruct", func(t *testing.T) {
			type userList struct {
				User  user
				Count uint8 `validate:"numeric"`
			}

			s := userList{
				User: user{
					"John",
					"Doe",
					42,
					"john.doe@domain.tld",
				},
				Count: 1,
			}

			structFields := []map[string]string{
				{"FirstName": "alpha"},
				{"LastName": "alpha"},
				{"Age": "number"},
				{"Email": "email"},
				{"Count": "numeric"},
			}

			fields := structWalker(s)
			for i, field := range fields {
				for key, value := range structFields[i] {
					assert.Equal(t, key, field.Name)
					assert.Equal(t, value, field.Tag)
				}
			}

			assert.Equal(t, len(structFields), len(fields))
		})
		t.Run("Slice", func(t *testing.T) {
			type userList struct {
				Users []user
			}

			s := userList{
				Users: []user{
					{
						"John",
						"Doe",
						42,
						"john.doe@domain.tld",
					},
					{
						"John",
						"Doe",
						42,
						"john.doe@domain.tld",
					},
				},
			}

			structFields := []map[string]string{
				{"FirstName": "alpha"},
				{"LastName": "alpha"},
				{"Age": "number"},
				{"Email": "email"},
				{"FirstName": "alpha"},
				{"LastName": "alpha"},
				{"Age": "number"},
				{"Email": "email"},
			}

			fields := structWalker(s)
			for i, field := range fields {
				for key, value := range structFields[i] {
					assert.Equal(t, key, field.Name)
					assert.Equal(t, value, field.Tag)
				}
			}

			assert.Equal(t, len(structFields), len(fields))
		})
		t.Run("UnexportedFields", func(t *testing.T) {
			type unexported struct {
				hidden string `validate:"alpha"`
			}

			u := unexported{hidden: "test"}
			results := structWalker(u)

			assert.Equal(t, 0, len(results))
		})
	})
}
