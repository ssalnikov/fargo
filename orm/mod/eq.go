package mod

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/query"
)

// Eq is equal '=' expression modificator for two fields
func Eq(f1 field.Mapper, f2 field.Mapper) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
