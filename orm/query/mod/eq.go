package mod

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/query"
)

// Eq modificator add equal condition
func Eq(f field.Field) Modifier {
	return func(q *query.Query) {
	}
}
