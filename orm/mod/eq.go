package mod

import (
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

// Eq is equal '=' expression modificator for two fields
func Eq(f1 model.Field, f2 model.Field) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
